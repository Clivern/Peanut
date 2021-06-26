// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markbates/pkger"
)

// Home controller
func Home(c *gin.Context) {
	index, err := pkger.Open("github.com/clivern/peanut:/web/dist/index.html")

	if err != nil {
		panic(err)
	}

	content, _ := ioutil.ReadAll(index)

	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write([]byte(content))
	return
}
