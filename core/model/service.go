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

// ServiceRecord type
type ServiceRecord struct {
	ID          string            `json:"id"`
	Service     string            `json:"service"`
	Version     string            `json:"version"`
	Configs     map[string]string `json:"configs"`
	DeleteAfter string            `json:"deleteAfter"`
	CreatedAt   int64             `json:"createdAt"`
	UpdatedAt   int64             `json:"updatedAt"`
}

// Service type
type Service struct {
	db driver.Database
}

// NewServiceStore creates a new instance
func NewServiceStore(db driver.Database) *Service {
	result := new(Service)
	result.db = db

	return result
}

// CreateRecord stores a service record
func (s *Service) CreateRecord(record ServiceRecord) error {
	record.CreatedAt = time.Now().Unix()
	record.UpdatedAt = time.Now().Unix()

	result, err := util.ConvertToJSON(record)

	if err != nil {
		return err
	}

	log.WithFields(log.Fields{
		"service_id": record.ID,
		"service":    record.Service,
	}).Debug("Create a service record")

	// store service record data
	err = s.db.Put(fmt.Sprintf(
		"%s/service/%s/s-data",
		viper.GetString("app.database.etcd.databaseName"),
		record.ID,
	), result)

	if err != nil {
		return err
	}

	return nil
}

// UpdateRecord updates a service record
func (s *Service) UpdateRecord(record ServiceRecord) error {
	record.UpdatedAt = time.Now().Unix()

	result, err := util.ConvertToJSON(record)

	if err != nil {
		return err
	}

	log.WithFields(log.Fields{
		"service_id": record.ID,
		"service":    record.Service,
	}).Debug("Update a service record")

	// store service record data
	err = s.db.Put(fmt.Sprintf(
		"%s/service/%s/s-data",
		viper.GetString("app.database.etcd.databaseName"),
		record.ID,
	), result)

	if err != nil {
		return err
	}

	return nil
}

// GetRecord gets service record data
func (s *Service) GetRecord(serviceID string) (ServiceRecord, error) {
	recordData := &ServiceRecord{}

	log.WithFields(log.Fields{
		"service_id": serviceID,
	}).Debug("Get a service record data")

	data, err := s.db.Get(fmt.Sprintf(
		"%s/service/%s/s-data",
		viper.GetString("app.database.etcd.databaseName"),
		serviceID,
	))

	if err != nil {
		return *recordData, err
	}

	for k, v := range data {
		// Check if it is the data key
		if strings.Contains(k, "/s-data") {
			err = util.LoadFromJSON(recordData, []byte(v))

			if err != nil {
				return *recordData, err
			}

			return *recordData, nil
		}
	}

	return *recordData, fmt.Errorf(
		"Unable to find service record with id: %s",
		serviceID,
	)
}

// GetRecords get services
func (s *Service) GetRecords() ([]ServiceRecord, error) {
	records := make([]ServiceRecord, 0)

	log.Debug("Get services")

	data, err := s.db.Get(fmt.Sprintf(
		"%s/service",
		viper.GetString("app.database.etcd.databaseName"),
	))

	if err != nil {
		return records, err
	}

	for k, v := range data {
		if strings.Contains(k, "/s-data") {
			recordData := &ServiceRecord{}

			err = util.LoadFromJSON(recordData, []byte(v))

			if err != nil {
				return records, err
			}

			records = append(records, *recordData)
		}
	}

	return records, nil
}

// DeleteRecord deletes a service record
func (s *Service) DeleteRecord(serviceID string) (bool, error) {

	log.WithFields(log.Fields{
		"service_id": serviceID,
	}).Debug("Delete a service record")

	count, err := s.db.Delete(fmt.Sprintf(
		"%s/service/%s",
		viper.GetString("app.database.etcd.databaseName"),
		serviceID,
	))

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
