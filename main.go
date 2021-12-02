package main

import (
	"flag"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"github/aiwantaozi/github-actions-demo/pkg/version"
	"os"
)

var (
	KubeConfig string
)

func main() {
	app := cli.NewApp()
	app.Name = "testy"
	app.Version = version.FriendlyVersion()
	app.Usage = "testy needs help!"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "kubeconfig",
			EnvVar:      "KUBECONFIG",
			Destination: &KubeConfig,
		},
	}
	app.Action = run

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

func run(c *cli.Context) {
	flag.Parse()

	logrus.Info("Starting controller")
}
