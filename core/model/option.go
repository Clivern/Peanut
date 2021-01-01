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

// Option type
type Option struct {
	db driver.Database
}

// OptionData struct
type OptionData struct {
	Key       string `json:"key"`
	Value     string `json:"value"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
}

// NewOptionStore creates a new instance
func NewOptionStore(db driver.Database) *Option {
	result := new(Option)
	result.db = db

	return result
}

// CreateOption stores an option
func (o *Option) CreateOption(option OptionData) error {

	option.CreatedAt = time.Now().Unix()
	option.UpdatedAt = time.Now().Unix()

	result, err := util.ConvertToJSON(option)

	if err != nil {
		return err
	}

	log.WithFields(log.Fields{
		"option_key": option.Key,
	}).Debug("Create an option")

	// store option data
	err = o.db.Put(fmt.Sprintf(
		"%s/option/%s/o-data",
		viper.GetString("app.database.etcd.databaseName"),
		option.Key,
	), result)

	if err != nil {
		return err
	}

	return nil
}

// UpdateOptionByKey updates an option by key
func (o *Option) UpdateOptionByKey(option OptionData) error {
	option.UpdatedAt = time.Now().Unix()

	result, err := util.ConvertToJSON(option)

	if err != nil {
		return err
	}

	log.WithFields(log.Fields{
		"option_key": option.Key,
	}).Debug("Update an option")

	// store option data
	err = o.db.Put(fmt.Sprintf(
		"%s/option/%s/o-data",
		viper.GetString("app.database.etcd.databaseName"),
		option.Key,
	), result)

	if err != nil {
		return err
	}

	return nil
}

// UpdateOptions update options
func (o *Option) UpdateOptions(options []OptionData) error {
	log.Debug("Update options")

	for _, option := range options {
		err := o.UpdateOptionByKey(option)

		if err != nil {
			return err
		}
	}

	return nil
}

// GetOptionByKey gets an option by a key
func (o *Option) GetOptionByKey(key string) (OptionData, error) {
	optionResult := &OptionData{}

	log.WithFields(log.Fields{
		"option_key": key,
	}).Debug("Get an option")

	data, err := o.db.Get(fmt.Sprintf(
		"%s/option/%s/o-data",
		viper.GetString("app.database.etcd.databaseName"),
		key,
	))

	if err != nil {
		return *optionResult, err
	}

	for k, v := range data {
		// Check if it is the data key
		if strings.Contains(k, "/o-data") {
			err = util.LoadFromJSON(optionResult, []byte(v))

			if err != nil {
				return *optionResult, err
			}

			return *optionResult, nil
		}
	}

	return *optionResult, fmt.Errorf(
		"Unable to find an option with a key: %s",
		key,
	)
}

// DeleteOptionByKey deletes an option by a key
func (o *Option) DeleteOptionByKey(key string) (bool, error) {

	log.WithFields(log.Fields{
		"option_key": key,
	}).Debug("Delete an option")

	count, err := o.db.Delete(fmt.Sprintf(
		"%s/option/%s",
		viper.GetString("app.database.etcd.databaseName"),
		key,
	))

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
