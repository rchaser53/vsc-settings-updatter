package main

import (
	"io/ioutil"
	"net/http"
	"os"
)

type Connect interface {
	Get(url string) (http.Response, error)
}

func PullSettigsJson(url string, destPath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return NetworkError{Msg: err.Error()}
	}

	byteArray, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return IoError{Msg: err.Error()}
	}

	file, err := os.Create(destPath)
	if err != nil {
		return IoError{Msg: err.Error()}
	}
	defer file.Close()
	file.Write(byteArray)

	return nil
}
