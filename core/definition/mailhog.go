// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package definition

import (
	"fmt"
)

const (
	// MailhogService const
	MailhogService = "mailhog"

	// MailhogSMTPPort const
	MailhogSMTPPort = "1025"

	// MailhogHTTPPort const
	MailhogHTTPPort = "8025"

	// MailhogDockerImage const
	MailhogDockerImage = "mailhog/mailhog"

	// MailhogDockerImageVersion const
	MailhogDockerImageVersion = "v1.0.1"

	// MailhogRestartPolicy const
	MailhogRestartPolicy = "unless-stopped"
)

// GetMailhogConfig gets yaml definition object
func GetMailhogConfig(name, version string) DockerComposeConfig {
	services := make(map[string]Service)

	if version == "" {
		version = MailhogDockerImageVersion
	}

	services[name] = Service{
		Image:   fmt.Sprintf("%s:%s", MailhogDockerImage, version),
		Restart: MailhogRestartPolicy,
		Ports: []string{
			MailhogSMTPPort,
			MailhogHTTPPort,
		},
	}

	return DockerComposeConfig{
		Version:  "3",
		Services: services,
	}
}
