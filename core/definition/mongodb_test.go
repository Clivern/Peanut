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

// TestUnitMongoDB test cases
func TestUnitMongoDB(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("#TestMongoDB", func() {
		g.It("It should satisfy all provided test cases", func() {
			mongodb := GetMongoDBConfig("mongodb", "", "db", "user", "pass")
			result, err := mongodb.ToString()

			g.Assert(strings.Contains(result, fmt.Sprintf("image: %s", fmt.Sprintf("%s:%s", MongoDBDockerImage, MongoDBDockerImageVersion)))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf(`- "%s"`, MongoDBPort))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf("restart: %s", MongoDBRestartPolicy))).Equal(true)
			g.Assert(strings.Contains(result, "MONGO_INITDB_DATABASE=db")).Equal(true)
			g.Assert(strings.Contains(result, "MONGO_INITDB_ROOT_USERNAME=user")).Equal(true)
			g.Assert(strings.Contains(result, "MONGO_INITDB_ROOT_PASSWORD=pass")).Equal(true)
			g.Assert(err).Equal(nil)

			mongodb = GetMongoDBConfig("mongodb", "", "", "", "")
			result, err = mongodb.ToString()

			g.Assert(strings.Contains(result, fmt.Sprintf("image: %s", fmt.Sprintf("%s:%s", MongoDBDockerImage, MongoDBDockerImageVersion)))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf(`- "%s"`, MongoDBPort))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf("restart: %s", MongoDBRestartPolicy))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf("MONGO_INITDB_DATABASE=%s", MongoDBDefaultDatabase))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf("MONGO_INITDB_ROOT_USERNAME=%s", MongoDBDefaultUsername))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf("MONGO_INITDB_ROOT_PASSWORD=%s", MongoDBDefaultPassword))).Equal(true)
			g.Assert(err).Equal(nil)
		})
	})
}
