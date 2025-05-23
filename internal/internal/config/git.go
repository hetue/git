package config

import (
	"github.com/pangum/pangu"
)

type Git struct {
	*Credential `default:"{}" json:"credential,omitempty"`
	*Project    `default:"{}" json:"project,omitempty"`
	*Pull       `default:"{}" json:"pull,omitempty"`
	*Push       `default:"{}" json:"push,omitempty"`
	*Repository `default:"{}" json:"repository,omitempty"`

	Binary *Binary `default:"{}" json:"binary,omitempty"`
}

func newGit(config *pangu.Config) (git *Git, err error) {
	git = new(Git)
	err = config.Build().Get(git)

	return
}
