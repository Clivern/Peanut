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

// TestUnitConsul test cases
func TestUnitConsul(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("#TestConsul", func() {
		g.It("It should satisfy all provided test cases", func() {
			consul := GetConsulConfig("consul", "")
			result, err := consul.ToString()

			g.Assert(strings.Contains(
				result,
				fmt.Sprintf("image: %s", fmt.Sprintf("%s:%s", ConsulDockerImage, ConsulDockerImageVersion)),
			)).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf(`- "%s"`, ConsulHTTPPort))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf("restart: %s", ConsulRestartPolicy))).Equal(true)
			g.Assert(err).Equal(nil)
		})
	})
}
