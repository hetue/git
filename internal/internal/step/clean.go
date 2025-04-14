package step

import (
	"context"
	"os"
	"path/filepath"

	"github.com/hetue/core"
	"github.com/hetue/git/internal/internal/config"
	"github.com/hetue/git/internal/internal/step/internal/constant"
)

var _ core.Step = (*Clean)(nil)

type Clean struct {
	project *config.Project
}

func newClean(project *config.Project) *Clean {
	return &Clean{
		project: project,
	}
}

func (c *Clean) Name() string {
	return "清理"
}

func (c *Clean) Runnable() bool {
	return nil != c.project.Clear && *c.project.Clear
}

func (c *Clean) Retryable() bool { // 不重试
	return false
}

func (c *Clean) Asyncable() bool { // 不异步
	return false
}

func (c *Clean) Run(_ *context.Context) error {
	return os.RemoveAll(filepath.Join(c.project.Directory, constant.GitHome))
}
