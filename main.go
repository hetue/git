package main

import (
	"github.com/hetue/core"
	"github.com/hetue/git/internal"
)

func main() {
	bootstrap := core.New()
	bootstrap.Build().Boot(internal.New)
}
