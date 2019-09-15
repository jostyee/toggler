package release_test

import (
	"math/rand"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/adamluzsi/testcase"
	"github.com/toggler-io/toggler/services/release"
	. "github.com/toggler-io/toggler/testing"
)

func TestFeatureFlag(t *testing.T) {
	s := testcase.NewSpec(t)

	s.Let(`ReleaseFlagName`, func(t *testcase.T) interface{} { return ExampleName() })
	s.Let(`RolloutSeedSalt`, func(t *testcase.T) interface{} { return time.Now().Unix() })
	s.Let(`RolloutPercentage`, func(t *testcase.T) interface{} { return int(0) })
	s.Let(`RolloutApiURL`, func(t *testcase.T) interface{} { return nil })
	s.Let(`Flag`, func(t *testcase.T) interface{} {
		ff := &release.Flag{Name: t.I(`ReleaseFlagName`).(string)}
		ff.Rollout.RandSeed = t.I(`RolloutSeedSalt`).(int64)
		ff.Rollout.Strategy.Percentage = t.I(`RolloutPercentage`).(int)
		ff.Rollout.Strategy.DecisionLogicAPI = getRolloutApiURL(t)
		return ff
	})

	s.Describe(`Verify`, func(s *testcase.Spec) {
		subject := func(t *testcase.T) error { return getFeatureFlag(t).Verify() }

		s.When(`name is empty`, func(s *testcase.Spec) {
			s.Let(`ReleaseFlagName`, func(t *testcase.T) interface{} { return `` })

			s.Then(`error reported`, func(t *testcase.T) {
				require.Equal(t, release.ErrNameIsEmpty, subject(t))
			})
		})

		s.When(`Decision Logic API value`, func(s *testcase.Spec) {
			s.Context(`is an invalid request url`, func(s *testcase.Spec) {
				s.Context(`because it is empty`, func(s *testcase.Spec) {
					s.Let(`RolloutApiURL`, func(t *testcase.T) interface{} { return `` })

					s.Then(`error reported`, func(t *testcase.T) {
						require.Equal(t, release.ErrInvalidRequestURL, subject(t))
					})
				})

				s.Context(`because it is without schema`, func(s *testcase.Spec) {
					s.Let(`RolloutApiURL`, func(t *testcase.T) interface{} { return `/invalid/request/url` })

					s.Then(`error reported`, func(t *testcase.T) {
						require.Equal(t, release.ErrInvalidRequestURL, subject(t))
					})
				})

				s.Context(`because it lacks host`, func(s *testcase.Spec) {
					s.Let(`RolloutApiURL`, func(t *testcase.T) interface{} { return `http://:8080/asd` })

					s.Then(`error reported`, func(t *testcase.T) {
						require.Equal(t, release.ErrInvalidRequestURL, subject(t))
					})
				})
			})

			s.Context(`is nil`, func(s *testcase.Spec) {
				s.Let(`RolloutApiURL`, func(t *testcase.T) interface{} { return nil })

				s.Then(`accepted`, func(t *testcase.T) {
					require.Nil(t, subject(t))
				})
			})

			s.Context(`is a valid request url`, func(s *testcase.Spec) {
				s.Let(`RolloutApiURL`, func(t *testcase.T) interface{} { return `https://example.com` })

				s.Then(`accepted`, func(t *testcase.T) {
					require.Nil(t, subject(t))
				})
			})
		})

		s.When(`percentage`, func(s *testcase.Spec) {
			s.Context(`less than 0`, func(s *testcase.Spec) {
				s.Let(`RolloutPercentage`, func(t *testcase.T) interface{} { return -1 + (rand.Intn(1024) * -1) })

				s.Then(`it will report error regarding the percentage`, func(t *testcase.T) {
					require.Equal(t, release.ErrInvalidPercentage, subject(t))
				})
			})

			s.Context(`greater than 100`, func(s *testcase.Spec) {
				s.Let(`RolloutPercentage`, func(t *testcase.T) interface{} { return 101 + rand.Intn(1024) })

				s.Then(`it will report error regarding the percentage`, func(t *testcase.T) {
					require.Equal(t, release.ErrInvalidPercentage, subject(t))
				})
			})

			s.Context(`is a number between 0 and 100`, func(s *testcase.Spec) {
				s.Let(`RolloutPercentage`, func(t *testcase.T) interface{} { return rand.Intn(101) })

				s.Then(`accepted`, func(t *testcase.T) {
					require.Nil(t, subject(t))
				})
			})
		})
	})
}

func getFeatureFlag(t *testcase.T) *release.Flag {
	ff := t.I(`Flag`)

	if ff == nil {
		return nil
	}

	return ff.(*release.Flag)
}

func getRolloutApiURL(t *testcase.T) *url.URL {
	rurl := t.I(`RolloutApiURL`)

	if rurl == nil {
		return nil
	}

	u, err := url.Parse(rurl.(string))
	require.Nil(t, err)
	return u
}
