package cli

import (
	"strings"
	"url-shortener/app/services"

	"github.com/urfave/cli"
)

type CLIParams struct {
	URL       []string
	ShortKeys []string
}

func (p *CLIParams) GetURL() []string {
	return p.URL
}

func (p *CLIParams) GetShortKeys() []string {
	return p.ShortKeys
}

func shortenWrapper(c *cli.Context) error {
	cliURLParams := c.String("urls")
	urls := strings.Split(cliURLParams, ",")
	params := &CLIParams{URL: urls}
	shortURL, err := services.Shorten(params)
	if err != nil {
		return err
	}

	for _, url := range shortURL {
		println(url)
	}
	return nil
}

func redirectURLWrapper(c *cli.Context) error {
	url := c.String("keys")
	shortKeys := strings.Split(url, ",")
	params := &CLIParams{URL: nil, ShortKeys: shortKeys}
	shortURL, err := services.RedirectURL(params)
	if err != nil {
		return err
	}

	for _, url := range shortURL {
		println(url)
	}
	return nil
}

func Generate() *cli.App {
	app := cli.NewApp()

	shortenFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "urls",
			Value: "https://www.linkedin.com/in/leonardo-teixeira-c%C3%A2ndido-286065191/",
		},
	}

	redirectFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "keys",
			Value: "qsil3y",
		},
	}

	app.Name = "Command line application"
	app.Usage = "Generate shorten URLs"
	app.Commands = []cli.Command{
		{
			Name:   "shorten",
			Usage:  "Generate shorten URLs",
			Flags:  shortenFlags,
			Action: shortenWrapper,
		},
		{
			Name:   "key",
			Usage:  "Redirect to original URL",
			Flags:  redirectFlags,
			Action: redirectURLWrapper,
		},
	}

	return app
}
