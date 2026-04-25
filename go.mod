module git-downloader-tool

go 1.25.4

require (
	github.com/spf13/cobra v1.10.2
	github.com/spf13/pflag v1.0.10
	gopkg.in/yaml.v2 v2.4.0
)

require github.com/inconshreveable/mousetrap v1.1.0 // indirect

replace github.com/system66/git-downloader-tool => .
