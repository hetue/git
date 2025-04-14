package step

import (
	"context"
	"fmt"
	"os"

	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/log"
	"github.com/hetue/core"
	"github.com/hetue/git/internal/internal/config"
	"github.com/hetue/git/internal/internal/step/internal/constant"
)

var _ core.Step = (*Credential)(nil)

type Credential struct {
	runtime *core.Runtime
	config  *config.Credential
	logger  log.Logger
}

func newCredential(runtime *core.Runtime, credential *config.Credential, logger log.Logger) *Credential {
	return &Credential{
		runtime: runtime,
		config:  credential,
		logger:  logger,
	}
}

func (c *Credential) Name() string {
	return "授权"
}

func (c *Credential) Runnable() bool {
	return "" != c.config.Username && "" != c.config.Password
}

func (c *Credential) Retryable() bool { // 不重试
	return false
}

func (c *Credential) Asyncable() bool { // 不异步
	return false
}

func (c *Credential) Run(_ *context.Context) (err error) {
	filepath := c.runtime.Path(constant.NetrcFilename)
	if _, se := os.Stat(filepath); nil == se || os.IsExist(se) {
		_ = os.Remove(filepath)
	}

	content := fmt.Sprintf(constant.NetrcConfigFormatter, c.config.Username, c.config.Password)
	fields := gox.Fields[any]{
		field.New("filepath", filepath),
		field.New("username", c.config.Username),
		field.New("password", c.config.Password),
	}
	if err = os.WriteFile(filepath, []byte(content), constant.DefaultFilePerm); nil != err {
		c.logger.Error("写入授权文件出错", fields.Add(field.Error(err))...)
	} else {
		c.logger.Info("写入授权文件成功", fields...)
	}

	return
}
