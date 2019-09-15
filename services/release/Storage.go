package release

import (
	"context"
	"github.com/adamluzsi/frameless"
	"github.com/adamluzsi/frameless/resources"
)

type Storage interface {
	resources.Saver
	resources.Finder
	resources.Updater
	resources.Deleter
	resources.Truncater
	FlagFinder
	PilotFinder
}

type PilotEntries = frameless.Iterator
type FlagEntries = frameless.Iterator

type FlagFinder interface {
	FindReleaseFlagByName(ctx context.Context, name string) (*Flag, error)
	FindFlagsByName(ctx context.Context, names ...string) FlagEntries
}

type PilotFinder interface {
	FindReleaseFlagPilotByPilotExternalID(ctx context.Context, FeatureFlagID, ExternalPilotID string) (*Pilot, error)
	FindPilotsByFeatureFlag(ctx context.Context, ff *Flag) frameless.Iterator
	FindPilotEntriesByExtID(ctx context.Context, pilotExtID string) PilotEntries
}
