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
	PendingStatus = "@pending"

	// FailedStatus constant
	FailedStatus = "@failed"

	// SuccessStatus constant
	SuccessStatus = "@success"
)

// JobRecord type
type JobRecord struct {
	ID        string `json:"id"`
	Hostname  string `json:"hostname"`
	CronID    string `json:"cronId"`
	Status    string `json:"status"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
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

	result, err := util.ConvertToJSON(record)

	if err != nil {
		return err
	}

	log.WithFields(log.Fields{
		"job_id":   record.ID,
		"hostname": record.Hostname,
	}).Debug("Create a job record")

	// store job record data
	err = j.db.Put(fmt.Sprintf(
		"%s/host/%s/job/%s/j-data",
		viper.GetString("app.database.etcd.databaseName"),
		record.Hostname,
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
		"job_id":   record.ID,
		"hostname": record.Hostname,
	}).Debug("Update a job record")

	// store job record data
	err = j.db.Put(fmt.Sprintf(
		"%s/host/%s/job/%s/j-data",
		viper.GetString("app.database.etcd.databaseName"),
		record.Hostname,
		record.ID,
	), result)

	if err != nil {
		return err
	}

	return nil
}

// GetRecord gets job record data
func (j *Job) GetRecord(hostname, jobID string) (*JobRecord, error) {
	recordData := &JobRecord{}

	log.WithFields(log.Fields{
		"job_id":   jobID,
		"hostname": hostname,
	}).Debug("Get a job record data")

	data, err := j.db.Get(fmt.Sprintf(
		"%s/host/%s/job/%s/j-data",
		viper.GetString("app.database.etcd.databaseName"),
		hostname,
		jobID,
	))

	if err != nil {
		return recordData, err
	}

	for k, v := range data {
		// Check if it is the data key
		if strings.Contains(k, "/j-data") {
			err = util.LoadFromJSON(recordData, []byte(v))

			if err != nil {
				return recordData, err
			}

			return recordData, nil
		}
	}

	return recordData, fmt.Errorf(
		"Unable to find job record with id: %s and hostname: %s",
		jobID,
		hostname,
	)
}

// DeleteRecord deletes a job record
func (j *Job) DeleteRecord(hostname, jobID string) (bool, error) {

	log.WithFields(log.Fields{
		"job_id":   jobID,
		"hostname": hostname,
	}).Debug("Delete a job record")

	count, err := j.db.Delete(fmt.Sprintf(
		"%s/host/%s/job/%s",
		viper.GetString("app.database.etcd.databaseName"),
		hostname,
		jobID,
	))

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// GetHostJobs get jobs for a host
func (j *Job) GetHostJobs(hostname string) ([]*JobRecord, error) {

	log.Debug("Get jobs to run")

	records := make([]*JobRecord, 0)

	data, err := j.db.Get(fmt.Sprintf(
		"%s/host/%s/job",
		viper.GetString("app.database.etcd.databaseName"),
		hostname,
	))

	if err != nil {
		return records, err
	}

	for k, v := range data {
		// Check if it is the data key
		if strings.Contains(k, "/j-data") {
			recordData := &JobRecord{}

			err = util.LoadFromJSON(recordData, []byte(v))

			if err != nil {
				return records, err
			}

			records = append(records, recordData)
		}
	}

	return records, nil
}

// CountHostJobs counts host jobs
func (j *Job) CountHostJobs(hostname, cronID, status string) (int, error) {

	log.Debug("Count host jobs")

	count := 0

	data, err := j.db.Get(fmt.Sprintf(
		"%s/host/%s/job",
		viper.GetString("app.database.etcd.databaseName"),
		hostname,
	))

	if err != nil {
		return count, err
	}

	for k, v := range data {
		// Check if it is the data key
		if strings.Contains(k, "/j-data") {
			recordData := &JobRecord{}

			err = util.LoadFromJSON(recordData, []byte(v))

			if err != nil {
				return count, err
			}

			if recordData.Status != status || cronID != recordData.CronID {
				continue
			}

			count++
		}
	}

	return count, nil
}
