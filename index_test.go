package main

import (
	"os"
	"testing"
	customError "vscSettingUpdatter/error"
)

type MockStruct struct{}

func (m MockStruct) Bool(key string) bool {
	return true
}

func (m MockStruct) String(key string) string {
	return "test"
}

func TestIOError(t *testing.T) {
	err := ExecCli(MockStruct{})

	switch err.(type) {
	case customError.IoError:
		os.Exit(0)
	default:
		t.Error("err should be IoError")
	}
}
