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

// TestUnitRabbitMQ test cases
func TestUnitRabbitMQ(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("#TestRabbitMQ", func() {
		g.It("It should satisfy all provided test cases", func() {
			rabbitmq := GetRabbitMQConfig("rabbitmq", "")
			result, err := rabbitmq.ToString()

			g.Assert(strings.Contains(result, fmt.Sprintf("image: %s", fmt.Sprintf("%s:%s", RabbitMQDockerImage, RabbitMQDockerImageVersion)))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf(`- "%s"`, RabbitMQAMQPPort))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf(`- "%s"`, RabbitMQDashboardPort))).Equal(true)
			g.Assert(strings.Contains(result, fmt.Sprintf("restart: %s", RabbitMQRestartPolicy))).Equal(true)
			g.Assert(err).Equal(nil)
		})
	})
}
