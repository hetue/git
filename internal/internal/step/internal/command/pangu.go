package command

import (
	"github.com/pangum/pangu"
)

func init() {
	pangu.New().Get().Dependency().Puts(
		NewGit,
	).Build().Apply()
}
