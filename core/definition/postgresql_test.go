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

// TestUnitPostgreSQL test cases
func TestUnitPostgreSQL(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("#TestPostgreSQL", func() {
		g.It("It should satisfy all provided test cases", func() {
			postgresql := GetPostgreSQLConfig("postgresql", "", "db", "username", "password")
			result, err := postgresql.ToString()

			g.Assert(strings.Contains(result, fmt.Sprintf("image: %s", fmt.Sprintf("%s:%s", PostgreSQLDockerImage, PostgreSQLDockerImageVersion)))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf(`- "%s"`, PostgreSQLPort))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf("restart: %s", PostgreSQLRestartPolicy))).Equal(true)
			g.Assert(strings.Contains(result, "POSTGRES_DB=db")).Equal(true)
			g.Assert(strings.Contains(result, "POSTGRES_USER=username")).Equal(true)
			g.Assert(strings.Contains(result, "POSTGRES_PASSWORD=password")).Equal(true)
			g.Assert(err).Equal(nil)

			postgresql = GetPostgreSQLConfig("postgresql", "", "", "", "")
			result, err = postgresql.ToString()

			g.Assert(strings.Contains(result, fmt.Sprintf("image: %s", fmt.Sprintf("%s:%s", PostgreSQLDockerImage, PostgreSQLDockerImageVersion)))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf(`- "%s"`, PostgreSQLPort))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf("restart: %s", PostgreSQLRestartPolicy))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf("POSTGRES_DB=%s", PostgreSQLDefaultDatabase))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf("POSTGRES_USER=%s", PostgreSQLDefaultUsername))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf("POSTGRES_PASSWORD=%s", PostgreSQLDefaultPassword))).Equal(true)
			g.Assert(err).Equal(nil)
		})
	})
}
