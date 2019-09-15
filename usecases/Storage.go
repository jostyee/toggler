package usecases

import (
	"github.com/adamluzsi/frameless/resources"
	"github.com/toggler-io/toggler/services/release"
	"github.com/toggler-io/toggler/services/security"
	"io"
)

type Storage interface {
	resources.Saver
	resources.Finder
	resources.Updater
	resources.Deleter
	resources.Truncater
	release.FlagFinder
	release.PilotFinder
	security.TokenFinder
	io.Closer
}
