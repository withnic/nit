package main

import "testing"

func TestSearchGitDir(t *testing.T) {
	_, err := searchGitDir()
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
}
