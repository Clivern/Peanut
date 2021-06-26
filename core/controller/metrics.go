// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
)

var (
	workersCount = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "peanut",
			Name:      "workers_count",
			Help:      "Number of Async Workers",
		})

	queueCapacity = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "peanut",
			Name:      "workers_queue_capacity",
			Help:      "The maximum number of messages queue can process",
		})
)

func init() {
	prometheus.MustRegister(workersCount)
	prometheus.MustRegister(queueCapacity)
}

// Metrics controller
func Metrics() http.Handler {
	workersCount.Set(float64(viper.GetInt("app.broker.native.workers")))
	queueCapacity.Set(float64(viper.GetInt("app.broker.native.capacity")))

	return promhttp.Handler()
}
