package jconf

import (
	"testing"
)

type testConfig struct {
	MyString         string `json:"my_string"`
	MyInt            int
	MyDefaultString  string `default:"foo"`
	MyDefaultInt     int    `default:"77"`
	MyDefaultString1 string `json:"my_default_string1" default:"defaultstring"`
	MyDefaultString2 string `json:"i_do_not_exist" default:"defaultstring"`
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
	if config.MyDefaultString1 != "bar" {
		t.Error("JSON var MyDefaultString1 not correct, expected bar, got: ", config.MyDefaultString1)
	}
	if config.MyDefaultString2 != "defaultstring" {
		t.Error("JSON var MyDefaultString2 not correct, expected defaultstring, got: ", config.MyDefaultString2)
	}
}

func TestNoSuchFile(t *testing.T) {
	err := Load(&config, "iDontExist.json")
	if err == nil {
		t.Error("Test for non existing file failed")
	}
}
