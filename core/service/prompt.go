// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package service

import (
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
)

// Prompt struct
type Prompt struct {
}

// NotEmpty returns error if input is empty
func NotEmpty(input string) error {
	if strings.TrimSpace(input) == "" {
		return fmt.Errorf("Input must not be empty")
	}
	return nil
}

// Optional optional value
func Optional(_ string) error {
	return nil
}

// IsEmpty if field is empty
func IsEmpty(input string) bool {
	if strings.TrimSpace(input) == "" {
		return true
	}
	return false
}

// Input request a value from end user
func (p *Prompt) Input(label string, validate promptui.ValidateFunc) (string, error) {

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}

	item := promptui.Prompt{
		Label:     label,
		Templates: templates,
		Validate:  validate,
	}

	result, err := item.Run()

	if err != nil {
		return "", err
	}

	return result, nil
}

// Select request a value from a list from end user
func (p *Prompt) Select(label string, items []string) (string, error) {

	item := promptui.Select{
		Label: label,
		Items: items,
	}

	_, result, err := item.Run()

	if err != nil {
		return "", err
	}

	return result, nil
}
