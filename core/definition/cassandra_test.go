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

// TestUnitCassandra test cases
func TestUnitCassandra(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("#TestCassandra", func() {
		g.It("It should satisfy all provided test cases", func() {
			cassandra := GetCassandraConfig("cassandra", "")
			result, err := cassandra.ToString()

			g.Assert(strings.Contains(
				result,
				fmt.Sprintf("image: %s", fmt.Sprintf("%s:%s", CassandraDockerImage, CassandraDockerImageVersion)),
			)).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf(`- "%s"`, CassandraPort))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf("restart: %s", CassandraRestartPolicy))).Equal(true)
			g.Assert(err).Equal(nil)
		})
	})
}
