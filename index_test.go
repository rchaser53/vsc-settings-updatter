package main

import (
	"os"
	"testing"
	customError "vscSettingUpdatter/error"
)

type IOErrorMockStruct struct{}

func (m IOErrorMockStruct) Bool(key string) bool {
	return true
}

func (m IOErrorMockStruct) String(key string) string {
	return "test"
}

func TestIOError(t *testing.T) {
	err := ExecCli(IOErrorMockStruct{})

	switch err.(type) {
	case customError.IoError:
		os.Exit(0)
	default:
		t.Error("err should be IoError")
	}
}

type SamePathErrorMockStruct struct{}

func (m SamePathErrorMockStruct) Bool(key string) bool {
	return false
}

func (m SamePathErrorMockStruct) String(key string) string {
	return "test"
}

func TestSamePathError(t *testing.T) {
	err := ExecCli(SamePathErrorMockStruct{})

	switch err.(type) {
	case customError.SamePathError:
		os.Exit(0)
	default:
		t.Error("err should be SamePathError")
	}
}
