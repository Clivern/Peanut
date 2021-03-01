// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package model

import (
	"fmt"
	"strings"
	"time"

	"github.com/clivern/peanut/core/driver"
	"github.com/clivern/peanut/core/util"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	// PendingStatus constant
	PendingStatus = "PENDING"

	// FailedStatus constant
	FailedStatus = "FAILED"

	// SuccessStatus constant
	SuccessStatus = "SUCCESS"

	// DeployJob constant
	DeployJob = "@deployService"

	// DestroyJob constant
	DestroyJob = "@destroyService"
)

// JobRecord type
type JobRecord struct {
	ID        string        `json:"id"`
	Action    string        `json:"action"`
	Service   ServiceRecord `json:"service"`
	Status    string        `json:"status"`
	CreatedAt int64         `json:"createdAt"`
	UpdatedAt int64         `json:"updatedAt"`
}

// Job type
type Job struct {
	db driver.Database
}

// NewJobStore creates a new instance
func NewJobStore(db driver.Database) *Job {
	result := new(Job)
	result.db = db

	return result
}

// CreateRecord stores a job record
func (j *Job) CreateRecord(record JobRecord) error {
	record.CreatedAt = time.Now().Unix()
	record.UpdatedAt = time.Now().Unix()
	record.Service.CreatedAt = time.Now().Unix()
	record.Service.UpdatedAt = time.Now().Unix()

	result, err := util.ConvertToJSON(record)

	if err != nil {
		return err
	}

	log.WithFields(log.Fields{
		"job_id":     record.ID,
		"service_id": record.Service.ID,
	}).Debug("Create a job record")

	// store job record data
	err = j.db.Put(fmt.Sprintf(
		"%s/service/%s/job/%s/j-data",
		viper.GetString("app.database.etcd.databaseName"),
		record.Service.ID,
		record.ID,
	), result)

	if err != nil {
		return err
	}

	return nil
}

// UpdateRecord updates a job record
func (j *Job) UpdateRecord(record JobRecord) error {
	record.UpdatedAt = time.Now().Unix()

	result, err := util.ConvertToJSON(record)

	if err != nil {
		return err
	}

	log.WithFields(log.Fields{
		"job_id":     record.ID,
		"service_id": record.Service.ID,
	}).Debug("Update a job record")

	// store job record data
	err = j.db.Put(fmt.Sprintf(
		"%s/service/%s/job/%s/j-data",
		viper.GetString("app.database.etcd.databaseName"),
		record.Service.ID,
		record.ID,
	), result)

	if err != nil {
		return err
	}

	return nil
}

// GetRecord gets job record data
func (j *Job) GetRecord(serviceID, jobID string) (JobRecord, error) {
	recordData := &JobRecord{}

	log.WithFields(log.Fields{
		"job_id":     jobID,
		"service_id": serviceID,
	}).Debug("Get a job record data")

	data, err := j.db.Get(fmt.Sprintf(
		"%s/service/%s/job/%s/j-data",
		viper.GetString("app.database.etcd.databaseName"),
		serviceID,
		jobID,
	))

	if err != nil {
		return *recordData, err
	}

	for k, v := range data {
		// Check if it is the data key
		if strings.Contains(k, "/j-data") {
			err = util.LoadFromJSON(recordData, []byte(v))

			if err != nil {
				return *recordData, err
			}

			return *recordData, nil
		}
	}

	return *recordData, fmt.Errorf(
		"Unable to find job record with id: %s and service id: %s",
		jobID,
		serviceID,
	)
}

// DeleteRecord deletes a job record
func (j *Job) DeleteRecord(serviceID, jobID string) (bool, error) {

	log.WithFields(log.Fields{
		"job_id":     jobID,
		"service_id": serviceID,
	}).Debug("Delete a job record")

	count, err := j.db.Delete(fmt.Sprintf(
		"%s/service/%s/job/%s",
		viper.GetString("app.database.etcd.databaseName"),
		serviceID,
		jobID,
	))

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
