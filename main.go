package main

import (
	"strings"
	"url-shortener/app/config"
	"url-shortener/app/runner"
)

func main() {
	config.Load()

	var r runner.Runner

	if strings.ToUpper(config.Env) == "LOCAL" {
		r = &runner.CLIRunner{}
	} else {
		r = &runner.APIRunner{}
	}

	r.Run()
}
