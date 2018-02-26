package main

import (
	"fmt"

	"github.com/pkg/errors"
	gitwrapper "github.com/withnic/go-gitcmdwrapper"
)

type nit struct {
	configs []Config
}

// NewNit returns nit and error
func NewNit() (*nit, error) {
	cr := &ConfigReader{}
	gc, _ := cr.NewGlobalConfig()
	lc, _ := cr.NewLocalConfig()
	cfgs := make([]Config, 2)

	if gc != nil {
		cfgs = append(cfgs, *gc)
	}

	if lc != nil {
		cfgs = append(cfgs, *lc)
	}

	return &nit{
		configs: cfgs,
	}, nil
}

func (it *nit) Run(branch string, cmd string, args []string) error{
	switch cmd {
	case "push":
		if err := it.pushProc(branch, args); err != nil {
			return err
		}
	case "pull":
	case "commit":
	case "checkout":
	}


	return nil
}

func (it *nit) pushProc(branch string, args []string) error {
	if _, err := it.prePush(branch); err != nil {
		return err
	}

	if ret := gitwrapper.Exec(args); ret != 0 {
		return errors.New("Exec Error")
	}

	if _, err := it.afterPush(branch); err != nil {
		return err
	}

	return nil
}

// PrePush returns bool and error
func (it *nit) prePush(branch string) (bool, error) {
	if len(it.configs) == 0 {
		return true, nil
	}

	for _, config := range it.configs {
		for _, v := range config.Hooks {
			if len(v.PrePush.Forbiddens) > 0 {
				if !v.PrePush.canPush(branch) {
					err := fmt.Errorf("you can't allow push %s to %s", branch, branch)
					return false, err
				}
			}
		}
	}

	return true, nil
}

// AfterPush returns bool and error
func (it *nit) afterPush(branch string) (bool, error) {
	if len(it.configs) == 0 {
		return true, nil
	}


	return true, nil
}

func (it *PrePush) canPush(branch string) bool {
	for _, v := range it.Forbiddens {
		if v == branch {
			return false
		}
	}
	return true
}
