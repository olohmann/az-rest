package main

import (
	"github.com/jawher/mow.cli"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	app := cli.App("az-rest", "A simple Azure Resource Manager REST client.")
	app.Spec = "[-v]"

	var (
		verbose    = app.BoolOpt("v verbose", false, "Verbose output mode")
	)

	app.Before = func() {
		if *verbose {
			log.SetLevel(log.DebugLevel)
		} else {
			log.SetLevel(log.InfoLevel)
		}

	}

	app.Command("GET", "Issue a GET request", func(cmd *cli.Cmd) {
		var (
			apiVersion = cmd.StringOpt("a api-version", "", "The API version to use for each request. If not specified, latest version will be used.")
			rawUrl = cmd.StringArg("URL", "", "The URL to invoke.")
		)

		cmd.Spec = "-a URL"
		cmd.Action = func() {
			ArmGet(*rawUrl, *apiVersion, "")
		}
	})

	app.Command("POST", "Issue a POST request", func(cmd *cli.Cmd) {
		var (
			apiVersion = cmd.StringOpt("a api-version", "", "The API version to use for each request. If not specified, latest version will be used.")
			rawUrl  = cmd.StringArg("URL", "", "The URL to invoke.")
			reqBody = cmd.StringOpt("body", "", "The body for the POST request.")
		)

		cmd.Spec = "-a [--body] URL "
		cmd.Action = func() {
			ArmPost(*rawUrl, *apiVersion, "", *reqBody)
		}
	})

	app.Command("version", "Print version information", func(cmd *cli.Cmd) {
		cmd.Action = func() {
			log.Info("az-rest version: ", Version, " ", VersionPrerelease, " ", GitCommit)
		}
	})

	app.Run(os.Args)
}
