// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package service

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// FuzzyFinder type
type FuzzyFinder struct{}

// Available validates if fzf installed
func (f *FuzzyFinder) Available() bool {
	_, err := exec.LookPath("fzf")
	if err != nil {
		return false
	}

	return true
}

// Show shows a fzf list
func (f *FuzzyFinder) Show(items []string) (string, error) {
	shell := os.Getenv("SHELL")

	if len(shell) == 0 {
		shell = "sh"
	}

	cmd := exec.Command(shell, "-c", "fzf -m")
	cmd.Stderr = os.Stderr
	in, err := cmd.StdinPipe()

	if err != nil {
		return "", err
	}

	go func() {
		for _, item := range items {
			fmt.Fprintln(in, item)
		}
		in.Close()
	}()

	result, err := cmd.Output()

	if err != nil {
		return "", err
	}

	filtered := strings.Split(string(result), "\n")

	return filtered[0], nil
}
