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

// TestUnitElasticSearch test cases
func TestUnitElasticSearch(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("#TestElasticSearch", func() {
		g.It("It should satisfy all provided test cases", func() {
			elasticsearch := GetElasticSearchConfig("elasticsearch")
			result, err := elasticsearch.ToString()

			g.Assert(strings.Contains(result, fmt.Sprintf("image: %s", ElasticSearchDockerImage))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf(`- "%s"`, ElasticSearchRequestsPort))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf(`- "%s"`, ElasticSearchCommunicationPort))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf("restart: %s", ElasticSearchRestartPolicy))).Equal(true)
			g.Assert(strings.Contains(result, "discovery.type=single-node")).Equal(true)
			g.Assert(err).Equal(nil)
		})
	})
}
