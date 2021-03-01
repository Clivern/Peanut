// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package service

import (
	"context"
	"encoding/json"
	"fmt"
)

// DockerHubAPI const
const DockerHubAPI = "https://registry.hub.docker.com/v2/"

// DockerHub struct
type DockerHub struct {
	httpClient *HTTPClient
}

// TagsResponse type
type TagsResponse struct {
	Results []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"results"`
}

// LoadFromJSON update object from json
func (t *TagsResponse) LoadFromJSON(data []byte) error {
	err := json.Unmarshal(data, &t)

	if err != nil {
		return err
	}

	return nil
}

// ConvertToJSON convert object to json
func (t *TagsResponse) ConvertToJSON() (string, error) {
	data, err := json.Marshal(&t)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// NewDockerHub creates an instance of http client
func NewDockerHub(httpClient *HTTPClient) *DockerHub {
	return &DockerHub{
		httpClient: httpClient,
	}
}

// GetTags get image tags
func (d *DockerHub) GetTags(ctx context.Context, org, image string) ([]string, error) {
	var tags []string

	tagsResponse := &TagsResponse{}

	response, err := d.httpClient.Get(
		ctx,
		fmt.Sprintf("%s/repositories/%s/%s/tags/?page_size=1000", DockerHubAPI, org, image),
		map[string]string{},
		map[string]string{},
	)

	if err != nil {
		return tags, err
	}

	result, err := d.httpClient.ToString(response)

	if err != nil {
		return tags, err
	}

	err = tagsResponse.LoadFromJSON([]byte(result))

	if err != nil {
		return tags, err
	}

	for _, result := range tagsResponse.Results {
		tags = append(tags, result.Name)
	}

	return tags, nil
}
