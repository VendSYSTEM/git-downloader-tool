package config

func (c *Config) ApplyDefaults() {
    for i := range c.Repos {
        repo := &c.Repos[i]
        
        if repo.Remote == nil {
            repo.Remote = &c.Defaults.Remote
        }
        
        if repo.Revision == nil {
            repo.Revision = &c.Defaults.Revision
        }
        
        if repo.Path == nil {
            repo.Path = &c.Defaults.Path
        }
    }
}