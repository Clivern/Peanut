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

// TestUnitEtcd test cases
func TestUnitEtcd(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("#TestEtcd", func() {
		g.It("It should satisfy all provided test cases", func() {
			etcd := GetEtcdConfig("etcd")
			result, err := etcd.ToString()

			g.Assert(strings.Contains(result, fmt.Sprintf("image: %s", EtcdDockerImage))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf(`- "%s"`, EtcdPort))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf("restart: %s", EtcdRestartPolicy))).Equal(true)
			g.Assert(strings.Contains(result, "ALLOW_NONE_AUTHENTICATION=yes")).Equal(true)
			g.Assert(err).Equal(nil)
		})
	})
}
