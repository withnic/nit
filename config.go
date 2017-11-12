package main

// Config is config struct
type Config struct {
	Hooks []Hook `yaml:"hooks"`
}

// Hook is hook struct
type Hook struct {
	PrePush PrePush `yaml:"prepush"`
}

// PrePush is push config struct
type PrePush struct {
	Forbiddens []string `yaml:"forbidden"`
}
