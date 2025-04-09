package internal

import (
	"github.com/hetue/core"
	"github.com/hetue/git/internal/internal"
)

func New(params internal.Steps) []core.Step {
	return []core.Step{
		params.Clear,
		params.Credential,
		params.SSH,
		params.Pull,
		params.Push,
	}
}
