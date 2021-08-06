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

// TestUnitMinio test cases
func TestUnitMinio(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("#TestMinio", func() {
		g.It("It should satisfy all provided test cases", func() {
			minio := GetMinioConfig("minio", "", "", "")
			result, err := minio.ToString()

			g.Assert(strings.Contains(result, fmt.Sprintf("image: %s", fmt.Sprintf("%s:%s", MinioDockerImage, MinioDockerImageVersion)))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf(`- "%s"`, MinioAPIPort))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf(`- "%s"`, MinioConsolePort))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf("restart: %s", MinioRestartPolicy))).Equal(true)
			g.Assert(err).Equal(nil)
		})
	})
}
