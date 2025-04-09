package internal

import (
	"github.com/hetue/todo/internal/internal/step"
	"github.com/pangum/pangu"
)

type Steps struct {
	pangu.Get

	Todo *step.Todo
}
