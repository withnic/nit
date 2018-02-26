package main

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	yaml "gopkg.in/yaml.v1"

	homedir "github.com/mitchellh/go-homedir"
)

const (
	confName  = ".nit.yml" // need .git same dir or homedir
	recursive = 3          //find
)

type ConfigReader struct {
}

// NewLocalConfig return config by user dir
func (it *ConfigReader) NewGlobalConfig() (*Config, error) {
	dir, err := homedir.Dir()
	if err != nil {
		return nil, err
	}
	name, err := it.searchConfigFile(dir, 1)
	if err != nil {
		return nil, err
	}

	config, err := it.newConfig(name)
	if err != nil {
		return nil, err
	}
	return config, nil
}

// newConfig returns Config struct and error
func (it *ConfigReader) newConfig(name string) (*Config, error) {
	buf, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, err
	}

	var c Config
	err = yaml.Unmarshal(buf, &c)
	return &c, err
}

// NewLocalConfig returns config
func (it *ConfigReader) NewLocalConfig() (*Config, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	name, err := it.searchConfigFile(dir, recursive)
	if err != nil {
		return nil, err
	}
	config, _ := it.newConfig(name)
	if err != nil {
		return nil, err
	}
	return config, err
}

// find searches confName
func (it *ConfigReader) find(files []os.FileInfo) bool {
	for _, file := range files {
		if file.Name() == confName {
			return true
		}
	}
	return false
}

// recusiveFind returns path and error
func (it *ConfigReader) recusiveFind(dir string, times int) (string, error) {
	if times < 0 {
		return "", errors.New("Not found")
	}
	d := filepath.Dir(dir)
	files, _ := ioutil.ReadDir(d)
	if it.find(files) {
		return d, nil
	}

	return it.recusiveFind(d, times-1)
}

func (it *ConfigReader) searchDir(dir string, dep int) (string, error) {
	files, _ := ioutil.ReadDir(dir)
	if it.find(files) {
		return dir, nil
	}
	return it.recusiveFind(dir, dep)
}

// searchConfigFile
func (it *ConfigReader) searchConfigFile(dir string, dep int) (string, error) {
	dir, err := it.searchDir(dir, dep)
	if err != nil {
		return "", err
	}

	return filepath.Join(dir, confName), nil
}
