// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package model

import (
	"encoding/json"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	// COUNTER is a Prometheus COUNTER metric
	COUNTER string = "counter"
	// GAUGE is a Prometheus GAUGE metric
	GAUGE string = "gauge"
	// HISTOGRAM is a Prometheus HISTOGRAM metric
	HISTOGRAM string = "histogram"
	// SUMMARY is a Prometheus SUMMARY metric
	SUMMARY string = "summary"
)

// Metric struct
type Metric struct {
	Type    string            `json:"type"`
	Name    string            `json:"name"`
	Help    string            `json:"help"`
	Method  string            `json:"method"`
	Value   string            `json:"value"`
	Labels  prometheus.Labels `json:"labels"`
	Buckets []float64         `json:"buckets"`
}

// LoadFromJSON update object from json
func (m *Metric) LoadFromJSON(data []byte) (bool, error) {
	err := json.Unmarshal(data, &m)
	if err != nil {
		return false, err
	}
	return true, nil
}

// ConvertToJSON convert object to json
func (m *Metric) ConvertToJSON() (string, error) {
	data, err := json.Marshal(&m)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// LabelKeys gets a list of label keys
func (m *Metric) LabelKeys() []string {
	keys := []string{}

	for k := range m.Labels {
		keys = append(keys, k)
	}

	return keys
}

// LabelValues gets a list of label values
func (m *Metric) LabelValues() []string {
	values := []string{}

	for _, v := range m.Labels {
		values = append(values, v)
	}

	return values
}

// GetValueAsFloat gets a list of label values
func (m *Metric) GetValueAsFloat() (float64, error) {
	value, err := strconv.ParseFloat(m.Value, 64)

	if err != nil {
		return 0, nil
	}

	return value, nil
}
