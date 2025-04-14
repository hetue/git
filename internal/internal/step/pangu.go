package step

import (
	"github.com/pangum/pangu"
)

func init() {
	pangu.New().Get().Dependency().Puts(
		newClean,
		newCredential,
		newSSH,
		newPull,
		newPush,
	).Build().Apply()
}
