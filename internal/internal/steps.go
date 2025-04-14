package internal

import (
	"github.com/hetue/git/internal/internal/step"
	"github.com/pangum/pangu"
)

type Steps struct {
	pangu.Get

	Credential *step.Credential
	Pull       *step.Pull
	Push       *step.Push
	SSH        *step.SSH
}
