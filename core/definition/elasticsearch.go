// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package definition

import (
	"fmt"
)

const (
	// ElasticSearchService const
	ElasticSearchService = "elasticsearch"

	// ElasticSearchRequestsPort const
	ElasticSearchRequestsPort = "9200"

	// ElasticSearchCommunicationPort const
	ElasticSearchCommunicationPort = "9300"

	// ElasticSearchDockerImage const
	ElasticSearchDockerImage = "elasticsearch"

	// ElasticSearchDockerImageVersion const
	ElasticSearchDockerImageVersion = "7.13.3"

	// ElasticSearchRestartPolicy const
	ElasticSearchRestartPolicy = "unless-stopped"
)

// GetElasticSearchConfig gets yaml definition object
func GetElasticSearchConfig(name, version string) DockerComposeConfig {
	services := make(map[string]Service)

	if version == "" {
		version = ElasticSearchDockerImageVersion
	}

	services[name] = Service{
		Image:   fmt.Sprintf("%s:%s", ElasticSearchDockerImage, version),
		Restart: ElasticSearchRestartPolicy,
		Ports: []string{
			ElasticSearchRequestsPort,
			ElasticSearchCommunicationPort,
		},
		Environment: []string{"discovery.type=single-node"},
	}

	return DockerComposeConfig{
		Version:  "3",
		Services: services,
	}
}
