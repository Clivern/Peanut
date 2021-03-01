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

// TestUnitZipkin test cases
func TestUnitZipkin(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("#TestZipkin", func() {
		g.It("It should satisfy all provided test cases", func() {
			zipkin := GetZipkinConfig("zipkin", "")
			result, err := zipkin.ToString()

			g.Assert(strings.Contains(result, fmt.Sprintf("image: %s", fmt.Sprintf("%s:%s", ZipkinDockerImage, ZipkinDockerImageVersion)))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf(`- "%s"`, ZipkinPort))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf("restart: %s", ZipkinRestartPolicy))).Equal(true)
			g.Assert(err).Equal(nil)
		})
	})
}
