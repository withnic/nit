package main

// Config is config struct
type Config struct {
	Hooks []Hook `yaml:"hooks"`
}

// Hook is hook struct
type Hook struct {
	PrePush PrePush `yaml:"prepush"`
	AfterPush AfterPush `yaml:"afterpush"`
	PrePull PrePull `yaml:"prepull"`
}

type AfterPush struct {
	Actions []string `yaml:"forbidden"`
}

type PrePull struct {
	SameBranchOnly bool `yaml:"sameBranchOnly"`
}

// PrePush is push config struct
type PrePush struct {
	Forbiddens []string `yaml:"forbidden"`
}
