package main

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

const (
	confName   = ".nit.yml" // need .git same dir
	recursive  = 10
	searchName = ".git"
)

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

func (it *PrePush) canPush(branch string) bool {
	for _, v := range it.Forbiddens {
		if v == branch {
			return false
		}
	}
	return true
}

func find(files []os.FileInfo) bool {
	for _, file := range files {
		if file.IsDir() && file.Name() == searchName {
			return true
		}
	}
	return false
}

func recusiveFind(dir string, times int) (string, error) {
	if times < 0 {
		return "", errors.New("Not found")
	}
	d := filepath.Dir(dir)
	files, _ := ioutil.ReadDir(d)
	if find(files) {
		return d, nil
	}

	return recusiveFind(d, times-1)

}

func searchDir() (string, error) {
	dir, _ := os.Getwd()
	files, _ := ioutil.ReadDir(dir)
	if find(files) {
		return dir, nil
	}
	return recusiveFind(dir, 3)
}

func searchConfigFile() (string, error) {
	dir, err := searchDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(dir, confName), nil
}

// NewConfig returns Config struct and error
func NewConfig() (*Config, error) {
	name, err := searchConfigFile()
	if err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, err
	}

	var c Config
	err = yaml.Unmarshal(buf, &c)
	if err != nil {
		panic(err)
	}
	return &c, nil
}
