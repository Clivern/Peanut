// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/clivern/peanut/core/controller"
	"github.com/clivern/peanut/core/middleware"
	"github.com/clivern/peanut/core/util"

	"github.com/drone/envsubst"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var towerCmd = &cobra.Command{
	Use:   "api",
	Short: "Start peanut api server",
	Run: func(cmd *cobra.Command, args []string) {
		configUnparsed, err := ioutil.ReadFile(config)

		if err != nil {
			panic(fmt.Sprintf(
				"Error while reading config file [%s]: %s",
				config,
				err.Error(),
			))
		}

		configParsed, err := envsubst.EvalEnv(string(configUnparsed))

		if err != nil {
			panic(fmt.Sprintf(
				"Error while parsing config file [%s]: %s",
				config,
				err.Error(),
			))
		}

		viper.SetConfigType("yaml")
		err = viper.ReadConfig(bytes.NewBufferString(configParsed))

		if err != nil {
			panic(fmt.Sprintf(
				"Error while loading configs [%s]: %s",
				config,
				err.Error(),
			))
		}

		viper.SetDefault("app.name", util.GenerateUUID4())

		if viper.GetString("app.log.output") != "stdout" {
			dir, _ := filepath.Split(viper.GetString("app.log.output"))

			if !util.DirExists(dir) {
				if _, err := util.EnsureDir(dir, 775); err != nil {
					panic(fmt.Sprintf(
						"Directory [%s] creation failed with error: %s",
						dir,
						err.Error(),
					))
				}
			}

			if !util.FileExists(viper.GetString("app.log.output")) {
				f, err := os.Create(viper.GetString("app.log.output"))
				if err != nil {
					panic(fmt.Sprintf(
						"Error while creating log file [%s]: %s",
						viper.GetString("app.log.output"),
						err.Error(),
					))
				}
				defer f.Close()
			}
		}

		storage := viper.GetString("app.storage.path")

		if !util.DirExists(storage) {
			if _, err := util.EnsureDir(storage, 0775); err != nil {
				panic(fmt.Sprintf(
					"Directory [%s] creation failed with error: %s",
					storage,
					err.Error(),
				))
			}
		}

		if viper.GetString("app.log.output") == "stdout" {
			gin.DefaultWriter = os.Stdout
			log.SetOutput(os.Stdout)
		} else {
			f, _ := os.OpenFile(
				viper.GetString("app.log.output"),
				os.O_APPEND|os.O_CREATE|os.O_WRONLY,
				0775,
			)

			gin.DefaultWriter = io.MultiWriter(f)
			log.SetOutput(f)
		}

		lvl := strings.ToLower(viper.GetString("app.log.level"))
		level, err := log.ParseLevel(lvl)

		if err != nil {
			level = log.InfoLevel
		}

		log.SetLevel(level)

		if viper.GetString("app.mode") == "prod" {
			gin.SetMode(gin.ReleaseMode)
			gin.DefaultWriter = ioutil.Discard
			gin.DisableConsoleColor()
		}

		if viper.GetString("app.log.format") == "json" {
			log.SetFormatter(&log.JSONFormatter{})
		} else {
			log.SetFormatter(&log.TextFormatter{})
		}

		r := gin.Default()
		workers := controller.NewWorkers()

		// Allow CORS only for development
		if viper.GetString("app.mode") == "dev" {
			r.Use(middleware.Cors())
		}

		r.Use(middleware.Correlation())
		r.Use(middleware.Logger())
		r.Use(middleware.Metric())
		r.Use(middleware.Auth())

		r.GET("/favicon.ico", func(c *gin.Context) {
			c.String(http.StatusNoContent, "")
		})

		r.GET("/_health", controller.Health)
		r.GET("/_ready", controller.Ready)

		r.GET(
			viper.GetString("app.metrics.prometheus.endpoint"),
			gin.WrapH(controller.Metrics()),
		)

		apiv1 := r.Group("/api/v1")
		{
			apiv1.POST("/service", func(c *gin.Context) {
				rawBody, err := c.GetRawData()

				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"status": "error",
						"error":  "Invalid request",
					})
					return
				}

				workers.DeployRequest(c, rawBody)
			})

			apiv1.DELETE("/service/:serviceId", func(c *gin.Context) {
				workers.DestroyRequest(c)
			})

			apiv1.GET("/service", controller.GetServices)
			apiv1.GET("/service/:serviceId", controller.GetService)
			apiv1.GET("/job/:serviceId/:jobId", controller.GetJob)
			apiv1.GET("/tag/:serviceType/:fromCache", controller.GetTags)
		}

		dist, err := fs.Sub(Static, "web/dist")

		if err != nil {
			panic(fmt.Sprintf(
				"Error while accessing dist files: %s",
				err.Error(),
			))
		}

		staticServer := http.StripPrefix("/", http.FileServer(http.FS(dist)))

		r.NoRoute(gin.WrapH(staticServer))

		go workers.Watch()
		go workers.Finalize(workers.HandleWorkload())

		var runerr error

		if viper.GetBool("app.tls.status") {
			runerr = r.RunTLS(
				fmt.Sprintf(":%s", strconv.Itoa(viper.GetInt("app.port"))),
				viper.GetString("app.tls.pemPath"),
				viper.GetString("app.tls.keyPath"),
			)
		} else {
			runerr = r.Run(
				fmt.Sprintf(":%s", strconv.Itoa(viper.GetInt("app.port"))),
			)
		}

		if runerr != nil {
			panic(runerr.Error())
		}
	},
}

func init() {
	towerCmd.Flags().StringVarP(
		&config,
		"config",
		"c",
		"config.prod.yml",
		"Absolute path to config file (required)",
	)
	towerCmd.MarkFlagRequired("config")
	rootCmd.AddCommand(towerCmd)
}
