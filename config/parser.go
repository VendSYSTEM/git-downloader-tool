package config

import (
    "io/ioutil"
    "gopkg.in/yaml.v2"
)

func ParseConfig(filename string) (*Config, error) {
    data, err := ioutil.ReadFile(filename)
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