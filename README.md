# Package jconf

jconf is a Golang library aimed to help parse project configuration files written in JSON.

For example (see test file), this code:

```
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
```

Will produce the following struct:
```
{
    MyString:        "bar"
    MyInt:           2
    MyDefaultString: "foo"
    MyDefaultInt:    77
}
```

From the following JSON file:

```
{
    "my_string": "bar",
    "MyInt": 2
}
```

# Default tags

Only string and integers are supported. Tags values must be between quoted, but if the field having a default tag is an integer, it will be converted.

# TODO

- Allow to pass a non existant JSON file (to only load defaults)
- Use a logger
- Read env variables in addition to json?
