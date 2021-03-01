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

// TestUnitMySQL test cases
func TestUnitMySQL(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("#TestMySQL", func() {
		g.It("It should satisfy all provided test cases", func() {
			mysql := GetMySQLConfig("mysql", "mysql1", "mysql2", "mysql3", "mysql4")
			result, err := mysql.ToString()

			g.Assert(strings.Contains(result, fmt.Sprintf("image: %s", MySQLDockerImage))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf(`- "%s"`, MySQLPort))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf("restart: %s", MySQLRestartPolicy))).Equal(true)
			g.Assert(strings.Contains(result, "MYSQL_ALLOW_EMPTY_PASSWORD=no")).Equal(true)
			g.Assert(strings.Contains(result, "MYSQL_ROOT_PASSWORD=mysql1")).Equal(true)
			g.Assert(strings.Contains(result, "MYSQL_DATABASE=mysql2")).Equal(true)
			g.Assert(strings.Contains(result, "MYSQL_USER=mysql3")).Equal(true)
			g.Assert(strings.Contains(result, "MYSQL_PASSWORD=mysql4")).Equal(true)
			g.Assert(err).Equal(nil)

			mysql = GetMySQLConfig("mysql", "", "", "", "")
			result, err = mysql.ToString()

			g.Assert(strings.Contains(result, fmt.Sprintf("image: %s", MySQLDockerImage))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf(`- "%s"`, MySQLPort))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf("restart: %s", MySQLRestartPolicy))).Equal(true)
			g.Assert(strings.Contains(result, "MYSQL_ALLOW_EMPTY_PASSWORD=no")).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf("MYSQL_ROOT_PASSWORD=%s", MySQLDefaultRootPassword))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf("MYSQL_DATABASE=%s", MySQLDefaultDatabase))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf("MYSQL_USER=%s", MySQLDefaultUsername))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf("MYSQL_PASSWORD=%s", MySQLDefaultPassword))).Equal(true)
			g.Assert(err).Equal(nil)
		})
	})
}
