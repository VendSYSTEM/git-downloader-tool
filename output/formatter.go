package output

import (
    "fmt"
    "strings"
    "git-downloader-tool/config"
)

func FormatRemotes(remotes map[string]config.Remote) string {
    var output strings.Builder
    output.WriteString("Remotes:\n")
    
    for name, remote := range remotes {
        output.WriteString(fmt.Sprintf("  %s: %s\n", name, remote.URL))
    }
    
    return output.String()
}

func FormatDefaults(defaults config.Defaults) string {
    return fmt.Sprintf("Defaults:\n  Remote: %s\n  Revision: %s\n  Path: %s\n", 
        defaults.Remote, defaults.Revision, defaults.Path)
}

func FormatRepos(repos []config.Repo) string {
    var output strings.Builder
    output.WriteString("Repositories:\n")
    
    for _, repo := range repos {
        output.WriteString(fmt.Sprintf("  - name: %s\n", repo.Name))
        if repo.Remote != nil {
            output.WriteString(fmt.Sprintf("    remote: %s\n", *repo.Remote))
        }
        if repo.Revision != nil {
            output.WriteString(fmt.Sprintf("    revision: %s\n", *repo.Revision))
        }
        if repo.Path != nil {
            output.WriteString(fmt.Sprintf("    path: %s\n", *repo.Path))
        }
    }
    
    return output.String()
}