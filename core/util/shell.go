// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package util

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
)

// Exec run a shell command
func Exec(command string) (string, string, error) {
	var outb, errb bytes.Buffer

	items := strings.Split(
		command,
		" ",
	)

	_, err := exec.LookPath(items[0])

	if err != nil {
		return outb.String(), errb.String(), err
	}

	commands := strings.Split(
		command,
		" ",
	)

	cmd := exec.Command(commands[0], commands[1:]...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = &outb
	cmd.Stderr = &errb

	err = cmd.Run()

	if err != nil {
		return outb.String(), errb.String(), err
	}

	return outb.String(), errb.String(), nil
}
