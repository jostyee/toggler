package httpapi_test

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"

	"github.com/adamluzsi/testcase"
	"github.com/toggler-io/toggler/services/release"
	"github.com/stretchr/testify/require"

	. "github.com/toggler-io/toggler/testing"
)

func TestServeMux_SetPilotEnrollmentForFeature(t *testing.T) {
	s := testcase.NewSpec(t)
	s.Parallel()

	subject := func(t *testcase.T) *httptest.ResponseRecorder {
		rr := httptest.NewRecorder()
		NewServeMux(t).ServeHTTP(rr, t.I(`request`).(*http.Request))
		return rr
	}

	SetupSpecCommonVariables(s)

	s.Let(`enrollment query value`, func(t *testcase.T) interface{} {
		return strconv.FormatBool(GetPilotEnrollment(t))
	})

	s.Let(`TokenString`, func(t *testcase.T) interface{} {
		tt, _ := CreateToken(t, `manager`)
		return tt
	})

	s.Before(func(t *testcase.T) {
		require.Nil(t, GetStorage(t).Save(context.TODO(), GetReleaseFlag(t)))
	})

	s.Let(`request`, func(t *testcase.T) interface{} {
		u, err := url.Parse(`/rollout/flag/set-enrollment-manually.json`)
		require.Nil(t, err)

		values := u.Query()
		values.Set(`token`, t.I(`TokenString`).(string))
		values.Set(`pilot.flagID`, GetReleaseFlag(t).ID)
		values.Set(`pilot.extID`, GetExternalPilotID(t))
		values.Set(`pilot.enrolled`, t.I(`enrollment query value`).(string))
		u.RawQuery = values.Encode()

		return httptest.NewRequest(http.MethodGet, u.String(), bytes.NewBuffer([]byte{}))
	})

	s.When(`invalid enrollment given`, func(s *testcase.Spec) {
		s.Let(`enrollment query value`, func(t *testcase.T) interface{} {
			return `invalid`
		})

		s.Then(`it will return bad request`, func(t *testcase.T) {
			r := subject(t)

			require.Equal(t, 400, r.Code)
		})
	})

	s.When(`invalid token given`, func(s *testcase.Spec) {
		s.Let(`TokenString`, func(t *testcase.T) interface{} {
			return `invalid`
		})

		s.Then(`it will return unauthorized`, func(t *testcase.T) {
			r := subject(t)

			require.Equal(t, 401, r.Code)
		})
	})

	s.When(`valid token provided`, func(s *testcase.Spec) {
		s.Let(`TokenString`, func(t *testcase.T) interface{} {
			return CreateSecurityTokenString(t)
		})

		s.Then(`call succeed`, func(t *testcase.T) {
			r := subject(t)
			require.Equal(t, 200, r.Code)
		})

		s.Then(`pilot user enrollment set in the system`, func(t *testcase.T) {
			r := subject(t)
			require.Equal(t, 200, r.Code)

			var resp struct{}
			IsJsonResponse(t, r, &resp)

			p, err := GetStorage(t).FindReleaseFlagPilotByPilotExternalID(context.Background(), FindStoredReleaseFlagByName(t).ID, GetExternalPilotID(t))
			require.Nil(t, err)
			require.NotNil(t, p)
			require.Equal(t, GetPilotEnrollment(t), p.Enrolled)
		})

		s.And(`flag id is not existing`, func(s *testcase.Spec) {
			s.Before(func(t *testcase.T) {
				require.Nil(t, GetStorage(t).DeleteByID(context.Background(), release.Flag{}, GetReleaseFlag(t).ID))
			})

			s.Then(`bad request replied`, func(t *testcase.T) {
				require.Equal(t, http.StatusBadRequest, subject(t).Code)
			})
		})
	})

}
