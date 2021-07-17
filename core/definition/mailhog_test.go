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

// TestUnitMailhog test cases
func TestUnitMailhog(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("#TestMailhog", func() {
		g.It("It should satisfy all provided test cases", func() {
			mailhog := GetMailhogConfig("mailhog", "")
			result, err := mailhog.ToString()

			g.Assert(strings.Contains(result, fmt.Sprintf("image: %s", fmt.Sprintf("%s:%s", MailhogDockerImage, MailhogDockerImageVersion)))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf(`- "%s"`, MailhogSMTPPort))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf(`- "%s"`, MailhogHTTPPort))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf("restart: %s", MailhogRestartPolicy))).Equal(true)
			g.Assert(err).Equal(nil)
		})
	})
}
