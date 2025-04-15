package config

type Repository struct {
	// 远程仓库地址
	Url string `default:"${GIT_HTTP_URL}" validate:"required,url" json:"url,omitempty"`
	// 签出代码
	Checkout string `default:"${COMMIT}" json:"checkout,omitempty"`
}

func newRepository(git *Git) *Repository {
	return git.Repository
}
