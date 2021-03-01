// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package definition

import (
	"fmt"
)

const (
	// GraphiteService const
	GraphiteService = "graphite"

	// GraphiteWebPort const
	GraphiteWebPort = "80"

	// GraphiteCarbonPort const
	GraphiteCarbonPort = "2003"

	// GraphiteCarbonPicklePort const
	GraphiteCarbonPicklePort = "2004"

	// GraphiteCarbonAggregatorPort const
	GraphiteCarbonAggregatorPort = "2023"

	// GraphiteCarbonAggregatorPicklePort const
	GraphiteCarbonAggregatorPicklePort = "2024"

	// GraphiteStatsdPort const
	GraphiteStatsdPort = "8125"

	// GraphiteStatsdAdminPort const
	GraphiteStatsdAdminPort = "8126"

	// GraphiteDockerImage const
	GraphiteDockerImage = "graphiteapp/graphite-statsd"

	// GraphiteDockerImageVersion const
	GraphiteDockerImageVersion = "1.1.7-11"

	// GraphiteRestartPolicy const
	GraphiteRestartPolicy = "unless-stopped"
)

// GetGraphiteConfig gets yaml definition object
func GetGraphiteConfig(name, version string) DockerComposeConfig {
	services := make(map[string]Service)

	if version == "" {
		version = GraphiteDockerImageVersion
	}

	services[name] = Service{
		Image:   fmt.Sprintf("%s:%s", GraphiteDockerImage, version),
		Restart: GraphiteRestartPolicy,
		Ports: []string{
			GraphiteWebPort,
			fmt.Sprintf("%s-%s", GraphiteCarbonPort, GraphiteCarbonPicklePort),
			fmt.Sprintf("%s-%s", GraphiteCarbonAggregatorPort, GraphiteCarbonAggregatorPicklePort),
			GraphiteStatsdPort,
			GraphiteStatsdAdminPort,
		},
	}

	return DockerComposeConfig{
		Version:  "3",
		Services: services,
	}
}
