// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

// Message interface
type Message interface {
	GetCorrelation() string
	GetService() string
	GetJob() string
	GetType() string
}

// DeployRequest type
type DeployRequest struct {
	JobID         string            `json:"jobId"`
	ServiceID     string            `json:"serviceId"`
	Template      string            `json:"template"`
	Configs       map[string]string `json:"configs"`
	DeleteAfter   string            `json:"deleteAfter"`
	Type          string            `json:"type"`
	CorrelationID string            `json:"correlationID"`
}

// DestroyRequest type
type DestroyRequest struct {
	JobID         string            `json:"jobId"`
	ServiceID     string            `json:"serviceId"`
	Template      string            `json:"template"`
	Configs       map[string]string `json:"configs"`
	DeleteAfter   string            `json:"deleteAfter"`
	Type          string            `json:"type"`
	CorrelationID string            `json:"correlationID"`
}

// GetCorrelation gets the correlation id
func (d DeployRequest) GetCorrelation() string {
	return d.CorrelationID
}

// GetService gets the service id
func (d DeployRequest) GetService() string {
	return d.ServiceID
}

// GetJob gets the job id
func (d DeployRequest) GetJob() string {
	return d.JobID
}

// GetType gets the job type
func (d DeployRequest) GetType() string {
	return d.Type
}

// GetCorrelation gets the correlation id
func (d DestroyRequest) GetCorrelation() string {
	return d.CorrelationID
}

// GetService gets the service id
func (d DestroyRequest) GetService() string {
	return d.ServiceID
}

// GetJob gets the job id
func (d DestroyRequest) GetJob() string {
	return d.JobID
}

// GetType gets the job type
func (d DestroyRequest) GetType() string {
	return d.Type
}
