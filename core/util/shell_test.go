// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package util

import (
	"fmt"
	"testing"

	"github.com/clivern/peanut/pkg"

	"github.com/franela/goblin"
)

// TestUnitShell
func TestUnitShell(t *testing.T) {
	g := goblin.Goblin(t)

	baseDir := pkg.GetBaseDir("cache")

	g.Describe("#TestExecFunc", func() {
		g.It("It should satisfy test cases", func() {
			stout, sterr, err := Exec("echo Hello World")

			g.Assert(err).Equal(nil)
			g.Assert(stout).Equal("Hello World\n")
			g.Assert(sterr).Equal("")

			stout, sterr, err = Exec("lo Hello World")

			g.Assert(stout).Equal("")
			g.Assert(sterr).Equal("")
			g.Assert(err.Error()).Equal(`exec: "lo": executable file not found in $PATH`)

			stout, sterr, err = Exec(fmt.Sprintf("bash %s/bin/test_script.sh", baseDir))

			g.Assert(stout).Equal("Hello World\nExample of error with line number and message\n")
			g.Assert(sterr).Equal("test_script.sh: 13: An error has occurred.\n")
			g.Assert(err.Error()).Equal("exit status 1")
		})
	})
}
