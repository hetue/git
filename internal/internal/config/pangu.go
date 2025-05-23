package config

import (
	"github.com/pangum/pangu"
)

func init() {
	pangu.New().Get().Dependency().Puts(
		newGit,

		newBinary,
		newCredential,
		newProject,
		newPull,
		newPush,
		newRepository,
	).Build().Apply()
}
