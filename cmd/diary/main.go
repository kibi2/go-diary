package main

import (
	"fmt"
	"os"

	"github.com/komem3/go-diary/cmd"
	_ "github.com/komem3/go-diary/statik"
)

func main() {
	command := cmd.NewRootCommand(cmd.NewInitCommand(), cmd.NewFormatCommand(), cmd.NewCommand())
	if err := command.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
