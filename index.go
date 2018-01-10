package main

import (
	"io"
	"os"
	"path/filepath"

	"github.com/urfave/cli"

	customError "vscSettingUpdatter/error"
)

var settingsPath = "settings.json"
var abridgementHomePath = "~" + settingsPath
var homePath = filepath.Join(os.Getenv("HOME"), settingsPath)

func main() {
	app := cli.NewApp()
	createOption(app)

	app.Action = func(c *cli.Context) error {
		var err error

		if c.Bool("u") {
			err = copyFile(settingsPath, homePath)
		} else if c.Bool("l") {
			err = copyFile(homePath, settingsPath)
		}

		if err != nil {
			return customError.IoError{Msg: err.Error()}
		}

		filePath := convertHomePath(c.String("src"))
		destPath := convertHomePath(c.String("dest"))

		if filePath == destPath {
			return customError.SamePathError{Msg: "filepath shouldn't be the same of destPath"}
		}

		copyFile(filePath, destPath)

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		println(err.Error())
	}
}

func convertHomePath(path string) string {
	if path == abridgementHomePath {
		return homePath
	}
	return path
}

func createOption(app *cli.App) {
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "update, u",
			Usage: "update settings.json at HOME",
		},
		cli.BoolFlag{
			Name:  "load, l",
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

	app.Name = "vsc-settings-updatter"
	app.Version = "0.0.1"
}

func copyFile(src string, dest string) error {
	file, err := os.Open(src)
	if err != nil {
		return err
	}
	defer file.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, file)
	if err != nil {
		return err
	}

	return nil
}
