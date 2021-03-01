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

// TestUnitMemcached test cases
func TestUnitMemcached(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("#TestMemcached", func() {
		g.It("It should satisfy all provided test cases", func() {
			memcached := GetMemcachedConfig("memcached", "")
			result, err := memcached.ToString()

			g.Assert(strings.Contains(result, fmt.Sprintf("image: %s", fmt.Sprintf("%s:%s", MemcachedDockerImage, MemcachedDockerImageVersion)))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf(`- "%s"`, MemcachedPort))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf("restart: %s", MemcachedRestartPolicy))).Equal(true)
			g.Assert(err).Equal(nil)
		})
	})
}
