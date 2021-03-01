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

// TestUnitJaeger test cases
func TestUnitJaeger(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("#TestJaeger", func() {
		g.It("It should satisfy all provided test cases", func() {
			jaeger := GetJaegerConfig("jaeger", "")
			result, err := jaeger.ToString()

			g.Assert(strings.Contains(result, fmt.Sprintf("image: %s", fmt.Sprintf("%s:%s", JaegerDockerImage, JaegerDockerImageVersion)))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf(`- "%s"`, JaegerUDPPort1))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf(`- "%s"`, JaegerUDPPort2))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf(`- "%s"`, JaegerUDPPort3))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf(`- "%s"`, JaegerHTTPPort1))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf(`- "%s"`, JaegerHTTPPort2))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf(`- "%s"`, JaegerHTTPPort3))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf(`- "%s"`, JaegerHTTPPort4))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf(`- "%s"`, JaegerHTTPPort5))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf("restart: %s", JaegerRestartPolicy))).Equal(true)
			g.Assert(err).Equal(nil)
		})
	})
}
