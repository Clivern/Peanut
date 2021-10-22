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

// TestUnitHttpBin test cases
func TestUnitHttpBin(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("#TestHttpBin", func() {
		g.It("It should satisfy all provided test cases", func() {
			httpbin := GetHttpbinConfig("httpbin", "")
			result, err := httpbin.ToString()

			g.Assert(strings.Contains(result, fmt.Sprintf("image: %s", fmt.Sprintf("%s:%s", HttpbinDockerImage, HttpbinDockerImageVersion)))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf(`- "%s"`, HttpbinPort))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf("restart: %s", HttpbinRestartPolicy))).Equal(true)
			g.Assert(err).Equal(nil)
		})
	})
}
