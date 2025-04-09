package get

import (
	"github.com/goexl/log"
	"github.com/hetue/git/internal/internal/config"
	"github.com/hetue/git/internal/internal/step/internal/command"
	"github.com/pangum/pangu"
)

type Push struct {
	pangu.Get

	Repository *config.Repository
	Project    *config.Project
	Credential *config.Credential
	Push       *config.Push

	Git    *command.Git
	Logger log.Logger
}
