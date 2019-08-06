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
	PathPartConfig string = "config"
	PathPartSecret string = "secret"
	PathPartDefault string = "default"
	PathPartOverride string = "override"
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
	var value string

	// These directories are weighted in order of precedence. Values specified
	// in multiple paths will be overridden by directories further the list.
	//
	// @see https://github.com/skpr/docs/blob/master/docs/config.md#deep-dive
	paths := []string{
		c.FilePath(PathPartConfig, PathPartDefault, key),
		c.FilePath(PathPartSecret, PathPartDefault, key),
		c.FilePath(PathPartConfig, PathPartOverride, key),
		c.FilePath(PathPartSecret, PathPartOverride, key),
	}

	for _, file := range paths {
		if _, err := os.Stat(file); os.IsNotExist(err) {
			continue
		}
		contents, err := ioutil.ReadFile(file)
		if err != nil {
			// This could indicate permissions issues with the mount, so exit.
			return "", err
		}
		value = string(contents)
	}

	return strings.TrimSuffix(value, DefaultTrimSuffix), nil
}

// DirPath returns the directory path of a specific config type and
// system/user provided value.
func(c *Config) DirPath(configType, defaultOrOverride string) string {
	return fmt.Sprintf("%s/%s/%s", c.Path, configType, defaultOrOverride)
}

// FilePath returns the file path of a specific config type, system/user
// provided value, and key.
func(c *Config) FilePath(configType, defaultOrOverride, key string) string {
	return fmt.Sprintf("%s/%s", c.DirPath(configType, defaultOrOverride), key)
}
