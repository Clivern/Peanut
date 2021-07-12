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

// TestUnitPrometheus test cases
func TestUnitPrometheus(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("#TestPrometheus", func() {
		g.It("It should satisfy all provided test cases", func() {
			prometheus := GetPrometheusConfig("prometheus", "/etc/peanut/storage/da2ce8ac-d33f-4dd9-a345-d76f2a4336be.yml")
			result, err := prometheus.ToString()

			g.Assert(strings.Contains(result, fmt.Sprintf("image: %s", PrometheusDockerImage))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf(`- "%s"`, PrometheusPort))).Equal(true)
			g.Assert(strings.Contains(
				result,
				"- /etc/peanut/storage/da2ce8ac-d33f-4dd9-a345-d76f2a4336be.yml:/etc/prometheus/prometheus.yml",
			)).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf("restart: %s", PrometheusRestartPolicy))).Equal(true)
			g.Assert(err).Equal(nil)
		})
	})
}
