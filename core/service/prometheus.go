// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package service

import (
	"fmt"

	"github.com/clivern/peanut/core/model"

	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

// Prometheus struct
type Prometheus struct{}

// NewPrometheus create a new instance of prometheus backend
func NewPrometheus() *Prometheus {
	return &Prometheus{}
}

// Send sends metrics to prometheus
func (p *Prometheus) Send(metrics []model.Metric) error {
	log.Info(fmt.Sprintf(
		"Send %d metrics to prometheus backend",
		len(metrics),
	))

	for _, metric := range metrics {
		switch metric.Type {
		case model.COUNTER:
			p.Counter(metric)

		case model.GAUGE:
			p.Gauge(metric)

		case model.HISTOGRAM:
			p.Histogram(metric)

		case model.SUMMARY:
			p.Summary(metric)

		default:
			return fmt.Errorf("metric with type %s not implemented yet", metric.Type)
		}
	}

	return nil
}

// Summary updates or creates a summary
func (p *Prometheus) Summary(item model.Metric) error {
	var metric prometheus.Summary

	value, _ := item.GetValueAsFloat()

	opts := prometheus.SummaryOpts{
		Name: item.Name,
		Help: item.Help,
	}
	if len(item.Labels) > 0 {
		vec := prometheus.NewSummaryVec(opts, item.LabelKeys())
		err := prometheus.Register(vec)
		if err != nil {
			if are, ok := err.(prometheus.AlreadyRegisteredError); ok {
				vec = are.ExistingCollector.(*prometheus.SummaryVec)
			} else {
				return err
			}
		}

		metric = vec.With(item.Labels).(prometheus.Summary)
	} else {
		metric = prometheus.NewSummary(opts)
		err := prometheus.Register(metric)
		if err != nil {
			if are, ok := err.(prometheus.AlreadyRegisteredError); ok {
				metric = are.ExistingCollector.(prometheus.Summary)
			} else {
				return err
			}
		}
	}

	if item.Method == "observe" {
		metric.Observe(value)
	} else {
		return fmt.Errorf("method %s is not implemented yet", item.Method)
	}

	return nil
}

// Counter updates or creates a counter
func (p *Prometheus) Counter(item model.Metric) error {
	var metric prometheus.Counter

	value, _ := item.GetValueAsFloat()

	opts := prometheus.CounterOpts{
		Name: item.Name,
		Help: item.Help,
	}

	if len(item.Labels) > 0 {
		vec := prometheus.NewCounterVec(opts, item.LabelKeys())

		err := prometheus.Register(vec)

		if err != nil {
			if are, ok := err.(prometheus.AlreadyRegisteredError); ok {
				vec = are.ExistingCollector.(*prometheus.CounterVec)
			} else {
				return err
			}
		}

		metric = vec.With(item.Labels)
	} else {
		metric = prometheus.NewCounter(opts)
		err := prometheus.Register(metric)
		if err != nil {
			if are, ok := err.(prometheus.AlreadyRegisteredError); ok {
				metric = are.ExistingCollector.(prometheus.Counter)
			} else {
				return err
			}
		}
	}

	switch item.Method {
	case "inc":
		metric.Inc()
	case "add":
		metric.Add(value)
	default:
		return fmt.Errorf("method %s is not implemented yet", item.Method)
	}

	return nil
}

// Histogram updates or creates a histogram
func (p *Prometheus) Histogram(item model.Metric) error {
	var metric prometheus.Histogram

	value, _ := item.GetValueAsFloat()

	opts := prometheus.HistogramOpts{
		Name:    item.Name,
		Help:    item.Help,
		Buckets: item.Buckets,
	}

	if len(item.Labels) > 0 {
		vec := prometheus.NewHistogramVec(opts, item.LabelKeys())
		err := prometheus.Register(vec)
		if err != nil {
			if are, ok := err.(prometheus.AlreadyRegisteredError); ok {
				vec = are.ExistingCollector.(*prometheus.HistogramVec)
			} else {
				return err
			}
		}

		metric = vec.With(item.Labels).(prometheus.Histogram)
	} else {
		metric = prometheus.NewHistogram(opts)
		err := prometheus.Register(metric)
		if err != nil {
			if are, ok := err.(prometheus.AlreadyRegisteredError); ok {
				metric = are.ExistingCollector.(prometheus.Histogram)
			} else {
				return err
			}
		}
	}

	if item.Method == "observe" {
		metric.Observe(value)
	} else {
		return fmt.Errorf("method %s is not implemented yet", item.Method)
	}

	return nil
}

// Gauge updates or creates a gauge
func (p *Prometheus) Gauge(item model.Metric) error {
	var metric prometheus.Gauge

	value, _ := item.GetValueAsFloat()

	opts := prometheus.GaugeOpts{
		Name: item.Name,
		Help: item.Help,
	}
	if len(item.Labels) > 0 {
		vec := prometheus.NewGaugeVec(opts, item.LabelKeys())
		err := prometheus.Register(vec)
		if err != nil {
			if are, ok := err.(prometheus.AlreadyRegisteredError); ok {
				vec = are.ExistingCollector.(*prometheus.GaugeVec)
			} else {
				return err
			}
		}

		metric = vec.With(item.Labels)
	} else {
		metric = prometheus.NewGauge(opts)
		err := prometheus.Register(metric)
		if err != nil {
			if are, ok := err.(prometheus.AlreadyRegisteredError); ok {
				metric = are.ExistingCollector.(prometheus.Gauge)
			} else {
				return err
			}
		}
	}

	switch item.Method {
	case "set":
		metric.Set(value)
	case "inc":
		metric.Inc()
	case "dec":
		metric.Dec()
	case "add":
		metric.Add(value)
	case "sub":
		metric.Sub(value)
	default:
		return fmt.Errorf("method %s is not implemented yet", item.Method)
	}

	return nil
}
