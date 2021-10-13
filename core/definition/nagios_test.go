// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package definition

import (
	"fmt"
	"strings"
	"testing"

	"github.com/franela/goblin"
)

// TestUnitNagios test cases
func TestUnitNagios(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("#TestNagios", func() {
		g.It("It should satisfy all provided test cases", func() {
			httpbin := GetNagiosConfig("httpbin", "")
			result, err := httpbin.ToString()

			g.Assert(strings.Contains(result, fmt.Sprintf("image: %s", fmt.Sprintf("%s:%s", NagiosDockerImage, NagiosDockerImageVersion)))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf(`- "%s"`, NagiosPort))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf("restart: %s", NagiosRestartPolicy))).Equal(true)
			g.Assert(err).Equal(nil)
		})
	})
}
