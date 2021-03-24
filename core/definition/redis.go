// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package definition

import (
	"fmt"
)

// GetRedis gets yaml definition object
func GetRedis(image, volume, port, restart string) *DockerCompose {
	services := make(map[string]Service)
	volumes := make(map[string]string)

	services["redis"] = Service{
		Image:   image,
		Restart: restart,
		Volumes: []string{
			fmt.Sprintf("%s:/data", volume),
		},
		Ports: []string{
			fmt.Sprintf("6379:%s", port),
		},
	}

	volumes[volume] = ""

	return &DockerCompose{
		Version:  "3",
		Services: services,
		Volumes:  volumes,
	}
}
