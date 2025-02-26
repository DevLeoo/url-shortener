package main

import (
	"strings"
	"url-shortener/app/config"
	redis "url-shortener/app/database"
	"url-shortener/app/runner"
)

func main() {
	config.Load()
	redis.Connect()

	var r runner.Runner

	if strings.ToUpper(config.Env) == "LOCAL" {
		r = &runner.CLIRunner{}
	} else {
		r = &runner.APIRunner{}
	}

	r.Run()
}
