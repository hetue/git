package command

import (
	"github.com/harluo/di"
)

func init() {
	di.New().Get().Dependency().Puts(
		NewGit,
	).Build().Apply()
}
