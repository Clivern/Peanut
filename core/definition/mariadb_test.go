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

// TestUnitMariaDB test cases
func TestUnitMariaDB(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("#TestMariaDB", func() {
		g.It("It should satisfy all provided test cases", func() {
			mariadb := GetMariaDBConfig("mariadb", "", "mariadb1", "mariadb2", "mariadb3", "mariadb4")
			result, err := mariadb.ToString()

			g.Assert(strings.Contains(result, fmt.Sprintf("image: %s", fmt.Sprintf("%s:%s", MariaDBDockerImage, MariaDBDockerImageVersion)))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf(`- "%s"`, MariaDBPort))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf("restart: %s", MariaDBRestartPolicy))).Equal(true)
			g.Assert(strings.Contains(result, "MYSQL_ALLOW_EMPTY_PASSWORD=no")).Equal(true)
			g.Assert(strings.Contains(result, "MYSQL_ROOT_PASSWORD=mariadb1")).Equal(true)
			g.Assert(strings.Contains(result, "MYSQL_DATABASE=mariadb2")).Equal(true)
			g.Assert(strings.Contains(result, "MYSQL_USER=mariadb3")).Equal(true)
			g.Assert(strings.Contains(result, "MYSQL_PASSWORD=mariadb4")).Equal(true)
			g.Assert(err).Equal(nil)

			mariadb = GetMariaDBConfig("mariadb", "", "", "", "", "")
			result, err = mariadb.ToString()

			g.Assert(strings.Contains(result, fmt.Sprintf("image: %s", fmt.Sprintf("%s:%s", MariaDBDockerImage, MariaDBDockerImageVersion)))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf(`- "%s"`, MariaDBPort))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf("restart: %s", MariaDBRestartPolicy))).Equal(true)
			g.Assert(strings.Contains(result, "MYSQL_ALLOW_EMPTY_PASSWORD=no")).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf("MYSQL_ROOT_PASSWORD=%s", MariaDBDefaultRootPassword))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf("MYSQL_DATABASE=%s", MariaDBDefaultDatabase))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf("MYSQL_USER=%s", MariaDBDefaultUsername))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf("MYSQL_PASSWORD=%s", MariaDBDefaultPassword))).Equal(true)
			g.Assert(err).Equal(nil)
		})
	})
}
