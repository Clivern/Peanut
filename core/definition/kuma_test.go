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

// TestUnitKuma test cases
func TestUnitKuma(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("#TestKuma", func() {
		g.It("It should satisfy all provided test cases", func() {
			kuma := GetKumaConfig("kuma", "")
			result, err := kuma.ToString()

			g.Assert(strings.Contains(
				result,
				fmt.Sprintf("image: %s", fmt.Sprintf("%s:%s", KumaDockerImage, KumaDockerImageVersion)),
			)).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf(`- "%s"`, KumaPort))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf("restart: %s", KumaRestartPolicy))).Equal(true)
			g.Assert(err).Equal(nil)
		})
	})
}
