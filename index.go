package main

import (
	"io"
	"os"
	"path/filepath"

	"github.com/urfave/cli"

	"vscSettingUpdatter/connect"
	customError "vscSettingUpdatter/error"
)

var settingsPath = "settings.json"
var abridgementHomePath = "~" + settingsPath
var homePath = filepath.Join(os.Getenv("HOME"), settingsPath)

type CliContext interface {
	Bool(key string) bool
	String(key string) string
}

func main() {
	app := cli.NewApp()
	createOption(app)

	app.Action = func(c *cli.Context) error {
		return ExecCli(c)
	}

	err := app.Run(os.Args)

	if err != nil {
		println(err.Error())
	}
}

func ExecCli(c CliContext) error {
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
	gitURL := convertHomePath(c.String("url"))

	if c.Bool("f") {
		fetchConfig := connect.TryGet(gitURL)
		file, err := os.Create(destPath)
		if err != nil {
			return customError.IoError{Msg: err.Error()}
		}

		defer file.Close()
		file.Write(([]byte)(fetchConfig))
	}

	if filePath == destPath {
		return customError.SamePathError{Msg: "filepath shouldn't be the same of destPath"}
	}

	copyFile(filePath, destPath)

	return nil
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
		cli.BoolFlag{
			Name:  "fetch, f",
			Usage: "fetch settings.json from github",
		},
		cli.StringFlag{
			Name:  "url",
			Usage: "github repository url",
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
