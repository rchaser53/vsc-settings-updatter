package main

import (
	"io"
	"os"
	"path/filepath"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "lang",
			Value: "english",
			Usage: "language for the greeting",
		},
	}

	app.Action = func(c *cli.Context) error {
		filePath := "settings.json"
		if c.NArg() > 0 {
			filePath = filepath.Join(c.Args().Get(0), "settings.json")
		}

		destPath := "settings.json"
		if c.NArg() > 1 {
			destPath = c.Args().Get(1)
		}

		copyFile(filePath, destPath)

		return nil
	}

	app.Run(os.Args)
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
