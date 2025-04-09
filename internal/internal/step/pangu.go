package step

import (
	"github.com/pangum/pangu"
)

func init() {
	pangu.New().Get().Dependency().Puts(
		newClear,
		newCredential,
		newSSH,
		newPull,
		newPush,
	).Build().Apply()
}
