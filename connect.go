package vscSettingUpdatter

import (
	"io/ioutil"
	"log"
	"net/http"
)

func TryGet(url string) string {
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	byteArray, _ := ioutil.ReadAll(resp.Body)
	bodyStr := string(byteArray)
	defer resp.Body.Close()

	return bodyStr
}
