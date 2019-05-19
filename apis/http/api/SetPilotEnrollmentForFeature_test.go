package api_test

import (
	"bytes"
	"github.com/adamluzsi/testcase"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"

	. "github.com/adamluzsi/FeatureFlags/testing"
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

	s.Let(`request`, func(t *testcase.T) interface{} {
		u, err := url.Parse(`/set-pilot-enrollment-for-feature`)
		require.Nil(t, err)

		values := u.Query()
		values.Set(`token`, t.I(`TokenString`).(string))
		values.Set(`feature`, GetFeatureFlagName(t))
		values.Set(`user-id`, GetExternalPilotID(t))
		values.Set(`enrollment`, strconv.FormatBool(GetPilotEnrollment(t)))
		u.RawQuery = values.Encode()

		return httptest.NewRequest(http.MethodGet, u.String(), bytes.NewBuffer([]byte{}))
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

			p, err := GetStorage(t).FindFlagPilotByExternalPilotID(FindFeatureFlag(t).ID, GetExternalPilotID(t))
			require.Nil(t, err)
			require.NotNil(t, p)
			require.Equal(t, GetPilotEnrollment(t), p.Enrolled)
		})
	})

}