package skprconfig

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const (
	DefaultPath string = "/etc/skpr"
	DefaultTrimSuffix string = "\n"
)

// Get returns the configured value of a given key, and the fallback value if no
// key does not exist.
//
// This is a convenience if the default skpr config mount point is being used.
func Get(key, fallback string) string {
	c := NewConfig(DefaultPath, DefaultTrimSuffix)
	return c.GetWithFallback(key, fallback)
}

// Config holds parameters for config.
type Config struct {
	Path string
	TrimSuffix string
}

// NewConfig returns a new Config struct with a given set of parameters.
func NewConfig(path, trimSuffix string) *Config {
	return &Config{
		Path: path,
		TrimSuffix: trimSuffix,
	}
}

// GetWithFallback returns the configured value of a given key, and the fallback
// value if no key does not exist.
func(c *Config) GetWithFallback(key, fallback string) string {
	value, err := c.Get(key)
	if err != nil {
		return fallback
	}

	return value
}

// Get returns the configured value of a given key.
func(c *Config) Get(key string) (string, error) {
	file := fmt.Sprintf("%s/%s", c.Path, key)

	if _, err := os.Stat(file); os.IsNotExist(err) {
		return "", err
	}

	contents, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}

	return strings.TrimSuffix(string(contents), "\n"), nil
}

