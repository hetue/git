package main

import (
	"github.com/hetue/core"
	"github.com/hetue/todo/internal"
)

func main() {
	bootstrap := core.New()
	bootstrap.Build().Boot(internal.New)
}
