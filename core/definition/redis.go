// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package definition

import (
	"fmt"
)

// Redis type
type Redis struct {
	Image   string   `yaml:"image"`
	Volumes []string `yaml:"volumes"`
	Ports   []string `yaml:"ports"`
	Restart string   `yaml:"restart"`
}

// GetRedis ..
func GetRedis(image, volume, port, restart string) *Compose {
	return &Compose{
		Version: "3",
		Services: Services{
			Redis: Redis{
				Image:   image,
				Restart: restart,
				Volumes: []string{
					fmt.Sprintf("%s:/data", volume),
				},
				Ports: []string{
					fmt.Sprintf("6379:%s", port),
				},
			},
		},
	}
}
