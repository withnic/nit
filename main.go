package main

import (
	"fmt"
	"os"
	"strings"

	gitwrapper "github.com/withnic/go-gitcmdwrapper"
)

var conf Config

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
	config, err := NewConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, " Error:%s\n", err.Error())
		return 1
	}

	cmd := strings.ToLower(args[0])
	branch := strings.ToLower(args[2])

	// push only
	if cmd != "push" {
		return gitwrapper.Exec(args)
	}

	for _, v := range config.Hooks {
		if len(v.PrePush.Forbiddens) > 0 {
			if !v.PrePush.canPush(branch) {
				err := fmt.Errorf("You Can't allow push %s to %s !!", branch, branch)
				fmt.Fprintf(os.Stderr, " Error:%s\n", err.Error())
				return 1
			}
		}
	}

	return gitwrapper.Exec(args)
}
