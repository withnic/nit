package main

import "fmt"

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

// CanPrePush returns bool and error
func (it *nit) CanPrePush(branch string) (bool, error) {
	if len(it.configs) == 0 {
		return true, nil
	}

	for _, config := range it.configs {
		for _, v := range config.Hooks {
			if len(v.PrePush.Forbiddens) > 0 {
				if !v.PrePush.canPush(branch) {
					err := fmt.Errorf("You Can't allow push %s to %s !!", branch, branch)
					return false, err
				}
			}
		}
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
