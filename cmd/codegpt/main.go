package main

import (
	"github.com/caoyuan-patsnap/CodeGPT/cmd"

	"github.com/caoyuan-patsnap/graceful"
)

func main() {
	m := graceful.NewManager()
	cmd.Execute(m.ShutdownContext())
}
