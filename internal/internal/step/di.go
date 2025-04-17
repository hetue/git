package step

import (
	"github.com/harluo/di"
)

func init() {
	di.New().Get().Dependency().Puts(
		newCredential,
		newSSH,
		newPull,
		newPush,
	).Build().Apply()
}
