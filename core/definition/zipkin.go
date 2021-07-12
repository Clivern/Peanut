// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package definition

const (
	// ZipkinService const
	ZipkinService = "zipkin"

	// ZipkinPort const
	ZipkinPort = "9411"

	// ZipkinDockerImage const
	ZipkinDockerImage = "openzipkin/zipkin:2.23"

	// ZipkinRestartPolicy const
	ZipkinRestartPolicy = "unless-stopped"
)

// GetZipkinConfig gets yaml definition object
func GetZipkinConfig(name string) DockerComposeConfig {
	services := make(map[string]Service)

	services[name] = Service{
		Image:   ZipkinDockerImage,
		Restart: ZipkinRestartPolicy,
		Ports:   []string{ZipkinPort},
	}

	return DockerComposeConfig{
		Version:  "3",
		Services: services,
	}
}
