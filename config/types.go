package config

// Remote represents a git remote configuration
type Remote struct {
	URL string `yaml:"url"`
}

// Defaults represents default values for repositories
type Defaults struct {
	Remote   string `yaml:"remote"`
	Revision string `yaml:"revision"`
	Path     string `yaml:"path"`
}

// Repository represents a git repository configuration
type Repository struct {
	Remote     string `yaml:"remote"`
	Repository string `yaml:"repository"`
	Revision   string `yaml:"revision"`
	Path       string `yaml:"path"`
}

// Config represents the overall configuration structure
type Config struct {
	Remotes  map[string]Remote     `yaml:"remotes"`
	Defaults Defaults              `yaml:"defaults"`
	Repos    map[string]Repository `yaml:"repos"`
}
