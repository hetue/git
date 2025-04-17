package config

import (
	"github.com/harluo/di"
)

func init() {
	di.New().Get().Dependency().Puts(
		newGit,

		newBinary,
		newCredential,
		newProject,
		newPull,
		newPush,
		newRepository,
	).Build().Apply()
}
