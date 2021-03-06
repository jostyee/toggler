package specs

import (
	"testing"

	frmlspecs "github.com/adamluzsi/frameless/resources/specs"
	rollspecs "github.com/toggler-io/toggler/services/release/specs"
	secuspecs "github.com/toggler-io/toggler/services/security/specs"

	"github.com/toggler-io/toggler/usecases"
)

type StorageSpec struct {
	Subject usecases.Storage
	frmlspecs.FixtureFactory
}

func (spec StorageSpec) Benchmark(b *testing.B) {
	b.Run(`toggler`, func(b *testing.B) {
		rollspecs.StorageSpec{Subject: spec.Subject, FixtureFactory: spec.FixtureFactory}.Benchmark(b)
		secuspecs.StorageSpec{Subject: spec.Subject, FixtureFactory: spec.FixtureFactory}.Benchmark(b)
	})
}

func (spec StorageSpec) Test(t *testing.T) {
	t.Run(`toggler`, func(t *testing.T) {
		rollspecs.StorageSpec{Subject: spec.Subject, FixtureFactory: spec.FixtureFactory}.Test(t)
		secuspecs.StorageSpec{Subject: spec.Subject, FixtureFactory: spec.FixtureFactory}.Test(t)
	})
}
