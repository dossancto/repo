package main

import (
	"os"

	"github.com/lu-css/repo/src/cli"
)

func main() {
	args := os.Args[1:]

	if len(args) >= 1 {
		cli.WithArgs(args)
		return
	}

	cli.Interative()
}
