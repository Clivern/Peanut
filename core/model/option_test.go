// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package model

import (
	"fmt"
	"testing"
	"time"

	"github.com/clivern/peanut/core/driver"
	"github.com/clivern/peanut/pkg"

	"github.com/franela/goblin"
	"github.com/spf13/viper"
)

// TestIntegrationOptionMethods test cases
func TestIntegrationOptionMethods(t *testing.T) {
	// Skip if -short flag exist
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	baseDir := pkg.GetBaseDir("cache")
	pkg.LoadConfigs(fmt.Sprintf("%s/config.dist.yml", baseDir))

	g := goblin.Goblin(t)

	db := driver.NewEtcdDriver()
	db.Connect()

	defer db.Close()

	option := NewOptionStore(db)

	// Cleanup
	db.Delete(viper.GetString("app.database.etcd.databaseName"))

	time.Sleep(3 * time.Second)

	g.Describe("#CreateOption", func() {
		g.It("It should satisfy test cases", func() {
			err := option.CreateOption(OptionData{
				Key:   "key_01",
				Value: "value_01",
			})

			g.Assert(err).Equal(nil)

			err = option.CreateOption(OptionData{
				Key:   "key_02",
				Value: "value_02",
			})

			g.Assert(err).Equal(nil)
		})
	})

	g.Describe("#UpdateOptionByKey", func() {
		g.It("It should satisfy test cases", func() {
			err := option.UpdateOptionByKey(OptionData{
				Key:       "key_02",
				Value:     "new_value_02",
				CreatedAt: time.Now().Unix(),
				UpdatedAt: time.Now().Unix(),
			})

			g.Assert(err).Equal(nil)

			err = option.UpdateOptionByKey(OptionData{
				Key:       "",
				Value:     "value_03",
				CreatedAt: time.Now().Unix(),
				UpdatedAt: time.Now().Unix(),
			})

			g.Assert(err).Equal(nil)
		})
	})

	g.Describe("#UpdateOptions", func() {
		g.It("It should satisfy test cases", func() {
			err := option.UpdateOptions([]OptionData{
				OptionData{
					Key:       "key_03",
					Value:     "new_value_03",
					CreatedAt: time.Now().Unix(),
					UpdatedAt: time.Now().Unix(),
				},
				OptionData{
					Key:       "key_04",
					Value:     "value_04",
					CreatedAt: time.Now().Unix(),
					UpdatedAt: time.Now().Unix(),
				},
			})

			g.Assert(err).Equal(nil)
		})
	})

	g.Describe("#GetOptionByKey", func() {
		g.It("It should satisfy test cases", func() {
			value, err := option.GetOptionByKey("key_03")

			g.Assert(value.Value).Equal("new_value_03")
			g.Assert(err).Equal(nil)
		})
	})

	g.Describe("#DeleteOptionByKey", func() {
		g.It("It should satisfy test cases", func() {
			ok, err := option.DeleteOptionByKey("key_04")

			g.Assert(ok).Equal(true)
			g.Assert(err).Equal(nil)
		})
	})
}
