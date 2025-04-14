package config

import (
	"strings"
)

type Repository struct {
	// 远程仓库地址
	Remote string `default:"${GIT_HTTP_URL}" validate:"required" json:"remote,omitempty"`
	// 分支
	Branch string `default:"master" json:"branch,omitempty"`
	// 提交
	Commit string `default:"${COMMIT}" validate:"required_without=Branch"`
}

func newRepository(git *Git) *Repository {
	return git.Repository
}

func (r *Repository) Checkout() (checkout string) {
	if "" != r.Commit {
		checkout = r.Commit
	} else {
		checkout = r.Branch
	}
	checkout = strings.TrimSpace(checkout)

	return
}
