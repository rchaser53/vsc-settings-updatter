package main

import (
	"io"
	"os"
	"log"
	"path/filepath"

	"github.com/urfave/cli"
)

var settingsPath = "settings.json"
var abridgementHomePath = "~" + settingsPath
var homePath = filepath.Join(os.Getenv("HOME"), settingsPath)

func main() {
	app := cli.NewApp()
	setupHelp(app)

	app.Action = func(c *cli.Context) error {
		if c.Bool("u") {
			updateSettingsJson()
			return nil
		} 

		if c.Bool("l") {
			loadSettingsJson()
			return nil
		}

		filePath := convertHomePath(c.String("src"))
		destPath := convertHomePath(c.String("dest"))

		if (filePath == destPath) {
			log.Fatal("filepath shouldn't be the same of destPath")
		}

		copyFile(filePath, destPath)

		return nil
	}

	app.Run(os.Args)
}

func updateSettingsJson() {
	copyFile(settingsPath, homePath)
}

func loadSettingsJson() {
	copyFile(homePath, settingsPath)
}

func convertHomePath(path string) string {
	if (path == abridgementHomePath) {
		return homePath
	}
	return path
}

func setupHelp(app *cli.App) {
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name: "update, u",
			Usage: "update settings.json at HOME",
		},
		cli.BoolFlag{
			Name: "load, l",
			Usage: "load settings.json at project",
		},
		cli.StringFlag{
			Name:  "src, s",
			Value: settingsPath,
			Usage: "imput path",
		},
		cli.StringFlag{
			Name:  "dest, d",
			Value: homePath,
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
