// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package service

import (
	"context"
	"net/http"
	"strings"
	"testing"

	"github.com/franela/goblin"
)

// TestUnitClientMethods
func TestUnitHttpClientMethods(t *testing.T) {
	g := goblin.Goblin(t)
	httpClient := NewHTTPClient(30)

	g.Describe("#Get", func() {
		g.It("It should satisfy test cases", func() {
			response, error := httpClient.Get(
				context.TODO(),
				"https://httpbin.org/get",
				map[string]string{"arg1": "value1"},
				map[string]string{"X-Api-Key": "poodle-123"},
			)

			g.Assert(httpClient.GetStatusCode(response)).Equal(http.StatusOK)
			g.Assert(error).Equal(nil)

			body, error := httpClient.ToString(response)

			g.Assert(strings.Contains(body, "value1")).Equal(true)
			g.Assert(strings.Contains(body, "arg1")).Equal(true)
			g.Assert(strings.Contains(body, "arg1=value1")).Equal(true)
			g.Assert(strings.Contains(body, "X-Api-Key")).Equal(true)
			g.Assert(strings.Contains(body, "poodle-123")).Equal(true)
			g.Assert(error).Equal(nil)
		})
	})

	g.Describe("#Delete", func() {
		g.It("It should satisfy test cases", func() {
			response, error := httpClient.Delete(
				context.TODO(),
				"https://httpbin.org/delete",
				map[string]string{"arg1": "value1"},
				map[string]string{"X-Api-Key": "poodle-123"},
			)

			g.Assert(httpClient.GetStatusCode(response)).Equal(http.StatusOK)
			g.Assert(error).Equal(nil)

			body, error := httpClient.ToString(response)

			g.Assert(strings.Contains(body, "value1")).Equal(true)
			g.Assert(strings.Contains(body, "arg1")).Equal(true)
			g.Assert(strings.Contains(body, "arg1=value1")).Equal(true)
			g.Assert(strings.Contains(body, "X-Api-Key")).Equal(true)
			g.Assert(strings.Contains(body, "poodle-123")).Equal(true)
			g.Assert(error).Equal(nil)
		})
	})

	g.Describe("#Post", func() {
		g.It("It should satisfy test cases", func() {
			response, error := httpClient.Post(
				context.TODO(),
				"https://httpbin.org/post",
				`{"Username":"admin", "Password":"12345"}`,
				map[string]string{"arg1": "value1"},
				map[string]string{"X-Api-Key": "poodle-123"},
			)

			g.Assert(httpClient.GetStatusCode(response)).Equal(http.StatusOK)
			g.Assert(error).Equal(nil)

			body, error := httpClient.ToString(response)

			g.Assert(strings.Contains(body, `"12345"`)).Equal(true)
			g.Assert(strings.Contains(body, `"Username"`)).Equal(true)
			g.Assert(strings.Contains(body, `"admin"`)).Equal(true)
			g.Assert(strings.Contains(body, `"Password"`)).Equal(true)
			g.Assert(strings.Contains(body, "value1")).Equal(true)
			g.Assert(strings.Contains(body, "arg1")).Equal(true)
			g.Assert(strings.Contains(body, "arg1=value1")).Equal(true)
			g.Assert(strings.Contains(body, "X-Api-Key")).Equal(true)
			g.Assert(strings.Contains(body, "poodle-123")).Equal(true)
			g.Assert(error).Equal(nil)
		})
	})

	g.Describe("#Put", func() {
		g.It("It should satisfy test cases", func() {
			response, error := httpClient.Put(
				context.TODO(),
				"https://httpbin.org/put",
				`{"Username":"admin", "Password":"12345"}`,
				map[string]string{"arg1": "value1"},
				map[string]string{"X-Api-Key": "poodle-123"},
			)

			g.Assert(httpClient.GetStatusCode(response)).Equal(http.StatusOK)
			g.Assert(error).Equal(nil)

			body, error := httpClient.ToString(response)

			g.Assert(strings.Contains(body, `"12345"`)).Equal(true)
			g.Assert(strings.Contains(body, `"Username"`)).Equal(true)
			g.Assert(strings.Contains(body, `"admin"`)).Equal(true)
			g.Assert(strings.Contains(body, `"Password"`)).Equal(true)
			g.Assert(strings.Contains(body, "value1")).Equal(true)
			g.Assert(strings.Contains(body, "arg1")).Equal(true)
			g.Assert(strings.Contains(body, "arg1=value1")).Equal(true)
			g.Assert(strings.Contains(body, "X-Api-Key")).Equal(true)
			g.Assert(strings.Contains(body, "poodle-123")).Equal(true)
			g.Assert(error).Equal(nil)
		})
	})

	g.Describe("#GetStatusCode", func() {
		g.It("It should satisfy test cases", func() {
			response, error := httpClient.Get(
				context.TODO(),
				"https://httpbin.org/status/200",
				map[string]string{"arg1": "value1"},
				map[string]string{"X-Api-Key": "poodle-123"},
			)

			g.Assert(httpClient.GetStatusCode(response)).Equal(http.StatusOK)
			g.Assert(error).Equal(nil)

			body, error := httpClient.ToString(response)

			g.Assert(body).Equal("")
			g.Assert(error).Equal(nil)
		})
	})

	g.Describe("#GetStatusCode", func() {
		g.It("It should satisfy test cases", func() {
			response, error := httpClient.Get(
				context.TODO(),
				"https://httpbin.org/status/500",
				map[string]string{"arg1": "value1"},
				map[string]string{"X-Api-Key": "poodle-123"},
			)

			g.Assert(httpClient.GetStatusCode(response)).Equal(http.StatusInternalServerError)
			g.Assert(error).Equal(nil)

			body, error := httpClient.ToString(response)

			g.Assert(body).Equal("")
			g.Assert(error).Equal(nil)
		})
	})

	g.Describe("#GetStatusCode", func() {
		g.It("It should satisfy test cases", func() {
			response, error := httpClient.Get(
				context.TODO(),
				"https://httpbin.org/status/404",
				map[string]string{"arg1": "value1"},
				map[string]string{"X-Api-Key": "poodle-123"},
			)

			g.Assert(httpClient.GetStatusCode(response)).Equal(http.StatusNotFound)
			g.Assert(error).Equal(nil)

			body, error := httpClient.ToString(response)

			g.Assert(body).Equal("")
			g.Assert(error).Equal(nil)
		})
	})

	g.Describe("#GetStatusCode", func() {
		g.It("It should satisfy test cases", func() {
			response, error := httpClient.Get(
				context.TODO(),
				"https://httpbin.org/status/201",
				map[string]string{"arg1": "value1"},
				map[string]string{"X-Api-Key": "poodle-123"},
			)

			g.Assert(httpClient.GetStatusCode(response)).Equal(http.StatusCreated)
			g.Assert(error).Equal(nil)

			body, error := httpClient.ToString(response)

			g.Assert(body).Equal("")
			g.Assert(error).Equal(nil)
		})
	})

	g.Describe("#buildParameters", func() {
		g.It("It should satisfy test cases", func() {
			url, error := httpClient.buildParameters("http://127.0.0.1", map[string]string{"arg1": "value1"})

			g.Assert(url).Equal("http://127.0.0.1?arg1=value1")
			g.Assert(error).Equal(nil)
		})
	})

	g.Describe("#BuildData", func() {
		g.It("It should satisfy test cases", func() {
			g.Assert(httpClient.BuildData(map[string]string{})).Equal("")
			g.Assert(httpClient.BuildData(map[string]string{"arg1": "value1"})).Equal("arg1=value1")
		})
	})
}
