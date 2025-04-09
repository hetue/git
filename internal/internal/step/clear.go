package step

import (
	"context"
	"os"
	"path/filepath"

	"github.com/hetue/core"
	"github.com/hetue/git/internal/internal/config"
	"github.com/hetue/git/internal/internal/step/internal/constant"
)

var _ core.Step = (*Clear)(nil)

type Clear struct {
	project *config.Project
}

func newClear(project *config.Project) *Clear {
	return &Clear{
		project: project,
	}
}

func (c *Clear) Name() string {
	return "清理"
}

func (c *Clear) Runnable() bool {
	return nil != c.project.Clear && *c.project.Clear
}

func (c *Clear) Retryable() bool { // 不重试
	return false
}

func (c *Clear) Asyncable() bool { // 不异步
	return false
}

func (c *Clear) Run(_ *context.Context) error {
	return os.RemoveAll(filepath.Join(c.project.Dir, constant.GitHome))
}
