package config

type Remote struct {
    Name string `yaml:"name"`
    URL  string `yaml:"url"`
}

type Defaults struct {
    Remote   string `yaml:"remote"`
    Revision string `yaml:"revision"`
    Path     string `yaml:"path"`
}

type Repo struct {
    Name     string  `yaml:"name"`
    Remote   *string `yaml:"remote,omitempty"`
    Revision *string `yaml:"revision,omitempty"`
    Path     *string `yaml:"path,omitempty"`
}

type Config struct {
    Remotes map[string]Remote `yaml:"remote"`
    Defaults Defaults          `yaml:"defaults"`
    Repos   []Repo            `yaml:"repos"`
}