package config

import (
	"github.com/pangum/pangu"
)

type Todo struct {
	// TODO
}

func newTodo(config *pangu.Config) (todo *Todo, err error) {
	todo = new(Todo)
	err = config.Build().Get(todo)

	return
}
