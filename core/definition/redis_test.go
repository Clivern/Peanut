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

// TestUnitRedis test cases
func TestUnitRedis(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("#TestRedis", func() {
		g.It("It should satisfy all provided test cases", func() {
			redis := GetRedisConfig("redis", "pass")
			result, err := redis.ToString()

			g.Assert(strings.Contains(result, fmt.Sprintf("image: %s", RedisDockerImage))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf(`- "%s"`, RedisPort))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf("restart: %s", RedisRestartPolicy))).Equal(true)
			g.Assert(strings.Contains(result, "REDIS_PASSWORD=pass")).Equal(true)
			g.Assert(err).Equal(nil)

			redis = GetRedisConfig("redis", "")
			result, err = redis.ToString()

			g.Assert(strings.Contains(result, fmt.Sprintf("image: %s", RedisDockerImage))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf(`- "%s"`, RedisPort))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf("restart: %s", RedisRestartPolicy))).Equal(true)
			g.Assert(strings.Contains(result, "ALLOW_EMPTY_PASSWORD=yes")).Equal(true)
			g.Assert(err).Equal(nil)
		})
	})
}
