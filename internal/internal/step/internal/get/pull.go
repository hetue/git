package get

import (
	"github.com/hetue/git/internal/internal/config"
	"github.com/hetue/git/internal/internal/step/internal/command"
	"github.com/pangum/pangu"
)

type Pull struct {
	pangu.Get

	Git        *command.Git
	Repository *config.Repository
	Project    *config.Project
	Credential *config.Credential
	Pull       *config.Pull
}
