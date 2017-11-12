package main

import (
	"fmt"
	"os"
	"strings"

	gitwrapper "github.com/withnic/go-gitcmdwrapper"
)

func main() {
	_, err := gitwrapper.Can()
	if err != nil {
		fmt.Fprintf(os.Stderr, " Error:%s\n", err.Error())
		os.Exit(1)
	}

	os.Exit(run(os.Args[1:]))
}

func run(args []string) int {
	// only support: nit push remote branch_name
	if len(args) != 3 {
		return gitwrapper.Exec(args)
	}

	cmd, branch := cmdBranchLower(args)
	// push only
	if cmd != "push" {
		return gitwrapper.Exec(args)
	}

	nit, err := NewNit()

	if nit == nil {
		return gitwrapper.Exec(args)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		return 1
	}

	if _, err := nit.CanPrePush(branch); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		return 1
	}

	return gitwrapper.Exec(args)
}

// cmdBranchLower returns cmd and branch
func cmdBranchLower(args []string) (string, string) {
	return strings.ToLower(args[0]), strings.ToLower(args[2])
}
