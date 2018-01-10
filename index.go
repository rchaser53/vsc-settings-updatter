package main

import (
	"io"
	"os"
	"log"
	"path/filepath"

	"github.com/urfave/cli"
)

var settingsPath = "settings.json"
var destDefaultPath = "~" + settingsPath

func main() {
	app := cli.NewApp()
	setupHelp(app)

	app.Action = func(c *cli.Context) error {
		filePath := c.String("src")
		destPath := c.String("dest")

		if destPath == destDefaultPath {
			destPath = filepath.Join(os.Getenv("HOME"), settingsPath)
		}

		if (filePath == destPath) {
			log.Fatal("filepath shouldn't be the same of destPath")
		}

		copyFile(filePath, destPath)

		return nil
	}

	app.Run(os.Args)
}

func setupHelp(app *cli.App) {
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "src, s",
			Value: settingsPath,
			Usage: "imput path",
		},
		cli.StringFlag{
			Name:  "dest, d",
			Value: destDefaultPath,
			Usage: "destination path",
		},
	}
}

func copyFile(src string, dest string) {
	file, err := os.Open(src)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		panic(err)
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, file)
	if err != nil {
		panic(err)
	}
}
