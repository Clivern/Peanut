// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package definition

import (
	"strings"
	"testing"

	"github.com/franela/goblin"
)

// TestUnitRedis test cases
func TestUnitRedis(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("#TestRedis", func() {
		g.It("It should satisfy all provided test cases", func() {
			redis := GetRedisConfig("redis:1.0.0", "redis_data", "1111", "always")
			result, err := redis.ToString()

			g.Assert(strings.Contains(result, "image: redis:1.0.0")).Equal(true)
			g.Assert(strings.Contains(result, "- redis_data:/data")).Equal(true)
			g.Assert(strings.Contains(result, "- 6379:1111")).Equal(true)
			g.Assert(strings.Contains(result, "restart: always")).Equal(true)
			g.Assert(strings.Contains(result, `redis_data: ""`)).Equal(true)
			g.Assert(err).Equal(nil)
		})
	})
}
