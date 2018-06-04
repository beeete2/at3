package main

import (
	"io"
	"flag"
	"fmt"
	"os"
)

const (
	VERSION = "0.0.1"
)

const (
	ExitCodeOK              int = 0
	ExitCodeError               = 10 + iota
	ExitCodeParseFlagsError
)

type CLI struct {
	outStream, errStream io.Writer
}

func (cli *CLI) Run(args []string) int {

	var (
		version bool
	)

	flags := flag.NewFlagSet("", flag.ContinueOnError)
	flags.SetOutput(cli.errStream)

	flags.BoolVar(&version, "version", false, "")
	flags.BoolVar(&version, "v", false, "")

	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeParseFlagsError
	}

	if version {
		fmt.Fprintf(cli.outStream, VERSION)
		return ExitCodeOK
	}

	at3 := NewAt3()
	err := at3.Transform(os.Stdin, os.Stdout)
	if err != nil {
		fmt.Fprintf(cli.errStream, "Log parse error: %s\n", err)
		return ExitCodeError
	}

	return ExitCodeOK
}
