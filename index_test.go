package main

import (
	"io/ioutil"
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
		return
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
		return
	default:
		t.Error("err should be SamePathError")
	}
}

type CopySettingsJsonMockStruct struct{}

func (m CopySettingsJsonMockStruct) Bool(key string) bool {
	return false
}

var srcPath = "fixtures/input/"
var comparisonPath = "fixtures/comparison"

func (m CopySettingsJsonMockStruct) String(key string) string {
	if key == "src" {
		return "fixtures/input/"
	}
	return "fixtures/comparison"
}
func TestCopySettingsJson(t *testing.T) {
	ExecCli(CopySettingsJsonMockStruct{})
	srcBytes, _ := ioutil.ReadFile("srcPath")
	destBytes, _ := ioutil.ReadFile("comparisonPath")

	if string(srcBytes) != string(destBytes) {
		t.Error("should copy fixtures/input/settings.json to fixtures/comparison/settings.json")
	}
}
