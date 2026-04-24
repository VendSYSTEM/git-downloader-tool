package main

import (
	"log"
	"git-downloader-tool/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}