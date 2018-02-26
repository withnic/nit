package main

import (
	"fmt"
	"os"
	"strings"

	"log"
	gitwrapper "github.com/withnic/go-gitcmdwrapper"
)

func main() {
	_, err := gitwrapper.Can()
	if err != nil {
		fmt.Fprintf(os.Stderr, " Error:%v\n", err)
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

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		return 1
	}

	if err = nit.Run(branch, cmd, args); err != nil{
		log.Fatal(err)
		return 1
	}

	return 0
}

// cmdBranchLower returns cmd and branch
func cmdBranchLower(args []string) (string, string) {
	return strings.ToLower(args[0]), strings.ToLower(args[2])
}
