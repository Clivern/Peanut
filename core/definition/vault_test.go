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

// TestUnitVault test cases
func TestUnitVault(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("#TestVault", func() {
		g.It("It should satisfy all provided test cases", func() {
			vault := GetVaultConfig("vault", "", "token")
			result, err := vault.ToString()

			g.Assert(strings.Contains(
				result,
				fmt.Sprintf("image: %s", fmt.Sprintf("%s:%s", VaultDockerImage, VaultDockerImageVersion)),
			)).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf(`- "%s"`, VaultHTTPPort))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf("restart: %s", VaultRestartPolicy))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf("VAULT_DEV_LISTEN_ADDRESS=0.0.0.0:%s", VaultHTTPPort))).Equal(true)
			g.Assert(strings.Contains(result, "vault server -dev -dev-root-token-id=token")).Equal(true)
			g.Assert(err).Equal(nil)
		})
	})
}
