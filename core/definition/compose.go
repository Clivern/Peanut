// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package definition

import (
	"gopkg.in/yaml.v2"
)

// Compose type
type Compose struct {
	Version  string   `yaml:"version"`
	Services Services `yaml:"services"`
}

// Services type
type Services struct {
	Redis Redis `yaml:"redis"`
}

// ToString converts type to string
func (c *Compose) ToString() (string, error) {
	d, err := yaml.Marshal(&c)

	if err != nil {
		return "", err
	}

	return string(d), nil
}
