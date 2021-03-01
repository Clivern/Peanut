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

// TestUnitGraphite test cases
func TestUnitGraphite(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("#TestGraphite", func() {
		g.It("It should satisfy all provided test cases", func() {
			graphite := GetGraphiteConfig("graphite", "")
			result, err := graphite.ToString()

			g.Assert(strings.Contains(result, fmt.Sprintf("image: %s", fmt.Sprintf("%s:%s", GraphiteDockerImage, GraphiteDockerImageVersion)))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf(`- "%s"`, GraphiteWebPort))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf(`- %s-%s`, GraphiteCarbonPort, GraphiteCarbonPicklePort))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf(`- %s-%s`, GraphiteCarbonAggregatorPort, GraphiteCarbonAggregatorPicklePort))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf(`- "%s"`, GraphiteStatsdPort))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf(`- "%s"`, GraphiteStatsdAdminPort))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf("restart: %s", GraphiteRestartPolicy))).Equal(true)
			g.Assert(err).Equal(nil)
		})
	})
}
