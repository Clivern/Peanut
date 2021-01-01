// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"

	"github.com/satori/go.uuid"
)

// InArray check if value is on array
func InArray(val interface{}, array interface{}) bool {
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) {
				return true
			}
		}
	}

	return false
}

// GenerateUUID4 create a UUID
func GenerateUUID4() string {
	u := uuid.Must(uuid.NewV4(), nil)
	return u.String()
}

// ReadFile get the file content
func ReadFile(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// FilterFiles filters files list based on specific sub-strings
func FilterFiles(files, filters []string) []string {
	var filteredFiles []string

	for _, file := range files {
		ok := true
		for _, filter := range filters {

			ok = ok && strings.Contains(file, filter)
		}
		if ok {
			filteredFiles = append(filteredFiles, file)
		}
	}

	return filteredFiles
}

// Unset remove element at position i
func Unset(a []string, i int) []string {
	a[i] = a[len(a)-1]
	a[len(a)-1] = ""
	return a[:len(a)-1]
}

// EnsureTrailingSlash ensure there is a trailing slash
func EnsureTrailingSlash(dir string) string {
	return fmt.Sprintf(
		"%s%s",
		strings.TrimRight(dir, string(os.PathSeparator)),
		string(os.PathSeparator),
	)
}

// RemoveTrailingSlash removes any trailing slash
func RemoveTrailingSlash(dir string) string {
	return strings.TrimRight(dir, string(os.PathSeparator))
}

// RemoveStartingSlash removes any starting slash
func RemoveStartingSlash(dir string) string {
	return strings.TrimLeft(dir, string(os.PathSeparator))
}

// ClearDir removes all files and sub dirs
func ClearDir(dir string) error {
	files, err := filepath.Glob(filepath.Join(dir, "*"))
	if err != nil {
		return err
	}
	for _, file := range files {
		err = os.RemoveAll(file)
		if err != nil {
			return err
		}
	}
	return nil
}

// StoreFile stores a file content
func StoreFile(path, content string) error {
	dir := filepath.Dir(path)

	err := os.MkdirAll(dir, 0775)

	if err != nil {
		return err
	}

	f, err := os.Create(path)

	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.WriteString(content)

	return err
}

// PathExists reports whether the path exists
func PathExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

// FileExists reports whether the named file exists
func FileExists(path string) bool {
	if fi, err := os.Stat(path); err == nil {
		if fi.Mode().IsRegular() {
			return true
		}
	}
	return false
}

// DirExists reports whether the dir exists
func DirExists(path string) bool {
	if fi, err := os.Stat(path); err == nil {
		if fi.Mode().IsDir() {
			return true
		}
	}
	return false
}

// EnsureDir ensures that directory exists
func EnsureDir(dirName string, mode int) (bool, error) {
	err := os.MkdirAll(dirName, os.FileMode(mode))

	if err == nil || os.IsExist(err) {
		return true, nil
	}
	return false, err
}

// DeleteFile deletes a file
func DeleteFile(path string) error {
	return os.Remove(path)
}

// LoadFromJSON update object from json
func LoadFromJSON(item interface{}, data []byte) error {
	err := json.Unmarshal(data, &item)
	if err != nil {
		return err
	}

	return nil
}

// ConvertToJSON convert object to json
func ConvertToJSON(item interface{}) (string, error) {
	data, err := json.Marshal(&item)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// GetHostname gets the hostname
func GetHostname() (string, error) {
	hostname, err := os.Hostname()

	if err != nil {
		return "", err
	}

	return strings.ToLower(hostname), nil
}

// DeleteDir deletes a dir
func DeleteDir(dir string) bool {
	if err := os.RemoveAll(dir); err == nil {
		return true
	}

	return false
}

// IsEmpty validate if string is empty or not
func IsEmpty(item string) bool {
	if strings.TrimSpace(item) == "" {
		return true
	}
	return false
}

// Rand gets a random number
func Rand(min, max int) int {
	return rand.Intn(max-min) + min
}

// GetVal gets a value from a hash map
func GetVal(hash map[string]string, key, def string) string {
	if val, ok := hash[key]; ok {
		return val
	}

	return def
}

// MergeMaps merges two maps
func MergeMaps(m1, m2 map[string]string) map[string]string {
	for k, v := range m2 {
		m1[k] = v
	}

	return m1
}

// TimeInSec time period in seconds
func TimeInSec(period string) int {
	if strings.HasSuffix(period, "sec") {
		i, _ := strconv.Atoi(strings.Replace(period, "sec", "", -1))
		return i
	} else if strings.HasSuffix(period, "min") {
		i, _ := strconv.Atoi(strings.Replace(period, "min", "", -1))
		return i * 60
	} else if strings.HasSuffix(period, "hours") {
		i, _ := strconv.Atoi(strings.Replace(period, "hours", "", -1))
		return i * 60 * 60
	} else if strings.HasSuffix(period, "days") {
		i, _ := strconv.Atoi(strings.Replace(period, "days", "", -1))
		return i * 60 * 60 * 24
	} else {
		return 0
	}
}
