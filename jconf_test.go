package jconf

import (
	"testing"
)

type testConfig struct {
	MyString        string `json:"my_string"`
	MyInt           int
	MyDefaultString string `default:"foo"`
	MyDefaultInt    int    `default:"77"`
}

var config testConfig

func init() {
	err := Load(&config, "testdata/test.json")
	if err != nil {
		panic(err)
	}
}

func TestLoad(t *testing.T) {
	if config.MyString != "bar" {
		t.Error("JSON var MyString not correct, expected bar, got: ", config.MyString)
	}
	if config.MyInt != 2 {
		t.Error("JSON var MyInt not correct, expected 2, got: ", config.MyInt)
	}
	if config.MyDefaultString != "foo" {
		t.Error("JSON var MyDefaultString not correct, expected foo, got: ", config.MyDefaultString)
	}
	if config.MyDefaultInt != 77 {
		t.Error("JSON var MyDefaultInt not correct, expected 77, got: ", config.MyDefaultInt)
	}
}

func TestNoSuchFile(t *testing.T) {
	err := Load(&config, "iDontExist.json")
	if err == nil {
		t.Error("Test for non existing file failed")
	}
}
