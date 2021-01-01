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
			redis := GetRedisConfig("redis", "redis:1.0.0", "6379", "always", "pass")
			result, err := redis.ToString()

			g.Assert(strings.Contains(result, "image: redis:1.0.0")).Equal(true)
			g.Assert(strings.Contains(result, `- "6379"`)).Equal(true)
			g.Assert(strings.Contains(result, "restart: always")).Equal(true)
			g.Assert(strings.Contains(result, "REDIS_PASSWORD=pass")).Equal(true)
			g.Assert(err).Equal(nil)

			redis = GetRedisConfig("redis", "redis:1.0.0", "6379", "always", "")
			result, err = redis.ToString()

			g.Assert(strings.Contains(result, "image: redis:1.0.0")).Equal(true)
			g.Assert(strings.Contains(result, `- "6379"`)).Equal(true)
			g.Assert(strings.Contains(result, "restart: always")).Equal(true)
			g.Assert(strings.Contains(result, "ALLOW_EMPTY_PASSWORD=yes")).Equal(true)
			g.Assert(err).Equal(nil)
		})
	})
}
