// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package service

import (
	"context"
	"testing"

	"github.com/franela/goblin"
)

// TestUnitGetTags test cases
func TestUnitGetTags(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("#GetTags", func() {
		g.It("It should satisfy all provided test cases", func() {
			gitHub := NewDockerHub(NewHTTPClient(20))
			tags, err := gitHub.GetTags(context.Background(), "library", "mysql")
			g.Assert(err).Equal(nil)
			g.Assert(tags[0] != "").Equal(true)
		})
	})
}
