// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package service

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// HTTPClient struct
type HTTPClient struct {
	Timeout time.Duration
}

// NewHTTPClient creates an instance of http client
func NewHTTPClient(timeout int) *HTTPClient {
	return &HTTPClient{
		Timeout: time.Duration(timeout),
	}
}

// Get http call
func (h *HTTPClient) Get(ctx context.Context, endpoint string, parameters, headers map[string]string) (*http.Response, error) {

	endpoint, err := h.buildParameters(endpoint, parameters)

	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest("GET", endpoint, nil)

	req = req.WithContext(ctx)

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	client := http.Client{
		Timeout: time.Second * h.Timeout,
	}

	resp, err := client.Do(req)

	if err != nil {
		return resp, err
	}

	return resp, err
}

// Post http call
func (h *HTTPClient) Post(ctx context.Context, endpoint string, data string, parameters, headers map[string]string) (*http.Response, error) {

	endpoint, err := h.buildParameters(endpoint, parameters)

	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest("POST", endpoint, bytes.NewBuffer([]byte(data)))

	req = req.WithContext(ctx)

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	client := http.Client{
		Timeout: time.Second * h.Timeout,
	}

	resp, err := client.Do(req)

	if err != nil {
		return resp, err
	}

	return resp, err
}

// Put http call
func (h *HTTPClient) Put(ctx context.Context, endpoint string, data string, parameters, headers map[string]string) (*http.Response, error) {

	endpoint, err := h.buildParameters(endpoint, parameters)

	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest("PUT", endpoint, bytes.NewBuffer([]byte(data)))

	req = req.WithContext(ctx)

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	client := http.Client{
		Timeout: time.Second * h.Timeout,
	}

	resp, err := client.Do(req)

	if err != nil {
		return resp, err
	}

	return resp, err
}

// Patch http call
func (h *HTTPClient) Patch(ctx context.Context, endpoint string, data string, parameters, headers map[string]string) (*http.Response, error) {

	endpoint, err := h.buildParameters(endpoint, parameters)

	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest("PATCH", endpoint, bytes.NewBuffer([]byte(data)))

	req = req.WithContext(ctx)

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	client := http.Client{
		Timeout: time.Second * h.Timeout,
	}

	resp, err := client.Do(req)

	if err != nil {
		return resp, err
	}

	return resp, err
}

// Delete http call
func (h *HTTPClient) Delete(ctx context.Context, endpoint string, parameters, headers map[string]string) (*http.Response, error) {

	endpoint, err := h.buildParameters(endpoint, parameters)

	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest("DELETE", endpoint, nil)

	req = req.WithContext(ctx)

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	client := http.Client{
		Timeout: time.Second * h.Timeout,
	}

	resp, err := client.Do(req)

	if err != nil {
		return resp, err
	}

	return resp, err
}

// buildParameters add parameters to URL
func (h *HTTPClient) buildParameters(endpoint string, parameters map[string]string) (string, error) {
	u, err := url.Parse(endpoint)

	if err != nil {
		return "", err
	}

	q := u.Query()

	for k, v := range parameters {
		q.Set(k, v)
	}

	u.RawQuery = q.Encode()

	return u.String(), nil
}

// BuildData build body data
func (h *HTTPClient) BuildData(parameters map[string]string) string {
	var items []string

	for k, v := range parameters {
		items = append(items, fmt.Sprintf("%s=%s", k, v))
	}

	return strings.Join(items, "&")
}

// ToString response body to string
func (h *HTTPClient) ToString(response *http.Response) (string, error) {
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}

// GetStatusCode response status code
func (h *HTTPClient) GetStatusCode(response *http.Response) int {
	return response.StatusCode
}

// GetHeaderValue get response header value
func (h *HTTPClient) GetHeaderValue(response *http.Response, key string) string {
	return response.Header.Get(key)
}
