package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// LoadConfig loads configuration from a file path
func LoadConfig(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

// MergeWithDefaults merges configuration with default values where needed
func (c *Config) MergeWithDefaults() error {
	// Validate and set defaults for each repository
	for name, repo := range c.Repos {
		if repo.Remote == "" {
			repo.Remote = c.Remotes[c.Defaults.Remote].URL
		}
		if repo.Revision == "" {
			repo.Revision = c.Defaults.Revision
		}
		if repo.Path == "" {
			repo.Path = c.Defaults.Path
		}
		c.Repos[name] = repo
	}
	return nil
}

// Validate checks that the configuration is valid
func (c *Config) Validate() error {
	// Validate remote URLs, repository names, paths, etc.
	if c.Defaults.Remote == "" {
		return fmt.Errorf("defaults.remote is required")
	}

	if _, exists := c.Remotes[c.Defaults.Remote]; !exists {
		return fmt.Errorf("default remote '%s' is not defined", c.Defaults.Remote)
	}

	for _, repo := range c.Repos {
		if repo.Repository == "" {
			return fmt.Errorf("repository name is required")
		}
	}

	return nil
}
