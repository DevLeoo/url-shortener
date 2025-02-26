package runner

import (
	"fmt"
	"log"
	"os"
	"url-shortener/app/cli"
)

type CLIRunner struct{}

func (c *CLIRunner) Run() {
	fmt.Println("Running CLI mode")

	aplicacao := cli.Generate()

	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
