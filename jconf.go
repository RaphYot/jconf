// jconf is a JSON configuration files parser for Go projects

package jconf

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

// Load will parse the JSON stored in a file `fileName` and update
// the `config` struct with that JSON config
func Load(config interface{}, fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("Opening config %s file failed: %v", fileName, err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return fmt.Errorf("Error decoding %s: %s", fileName, err)
	}
	// Let's compute the defaults tags and return errors if any
	return setDefaultsTags(config)
}

// setDefaultsTags will assign a default value to a field from its "default"
// tag in the struct. This will be set only if after parsing the field is still blank.
// Example:
//    MyDefaultString string `default:"foo"`
func setDefaultsTags(config interface{}) error {
	// Config should be a struct, reflecting it's type/value to verify the Tags
	configValue := reflect.Indirect(reflect.ValueOf(config))
	configType := configValue.Type()
	for i := 0; i < configType.NumField(); i++ {
		fieldStruct := configType.Field(i)
		field := configValue.Field(i)
		if isBlank := reflect.DeepEqual(field.Interface(), reflect.Zero(field.Type()).Interface()); isBlank {
			// set "default" tag if is blank
			if value := fieldStruct.Tag.Get("default"); value != "" {
				switch field.Kind() {
				case reflect.String:
					field.SetString(value)
				case reflect.Int:
					i, err := strconv.Atoi(value)
					if err != nil {
						return fmt.Errorf("Could not convert struct tag to integer: ", err)
					}
					field.SetInt(int64(i))
				default:
					return fmt.Errorf("This type is not supported for default tag: %v", field.Kind())
				}
			}
		}
	}
	return nil
}
