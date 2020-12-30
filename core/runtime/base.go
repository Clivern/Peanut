// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package runtime

import (
	"github.com/clivern/peanut/core/model"
)

const (
	// RedisDockerImage constant
	RedisDockerImage = "redis:5.0.10-alpine"
)

// Virtualization interface
type Virtualization interface {
	Deploy(service model.ServiceRecord) error
	Destroy(service model.ServiceRecord) error
}
