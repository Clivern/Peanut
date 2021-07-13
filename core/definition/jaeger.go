// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package definition

const (
	// JaegerService const
	JaegerService = "jaeger"

	// JaegerUDPPort1 const
	JaegerUDPPort1 = "5775"

	// JaegerUDPPort2 const
	JaegerUDPPort2 = "6831"

	// JaegerUDPPort3 const
	JaegerUDPPort3 = "6832"

	// JaegerHTTPPort1 const
	JaegerHTTPPort1 = "5778"

	// JaegerHTTPPort2 const
	JaegerHTTPPort2 = "16686"

	// JaegerHTTPPort3 const
	JaegerHTTPPort3 = "14268"

	// JaegerHTTPPort4 const
	JaegerHTTPPort4 = "14250"

	// JaegerHTTPPort5 const
	JaegerHTTPPort5 = "9411"

	// JaegerDockerImage const
	JaegerDockerImage = "jaegertracing/all-in-one:1.24"

	// JaegerRestartPolicy const
	JaegerRestartPolicy = "unless-stopped"
)

// GetJaegerConfig gets yaml definition object
func GetJaegerConfig(name string) DockerComposeConfig {
	services := make(map[string]Service)

	services[name] = Service{
		Image:   JaegerDockerImage,
		Restart: JaegerRestartPolicy,
		Ports: []string{
			JaegerUDPPort1,
			JaegerUDPPort2,
			JaegerUDPPort3,
			JaegerHTTPPort1,
			JaegerHTTPPort2,
			JaegerHTTPPort3,
			JaegerHTTPPort4,
			JaegerHTTPPort5,
		},
	}

	return DockerComposeConfig{
		Version:  "3",
		Services: services,
	}
}
