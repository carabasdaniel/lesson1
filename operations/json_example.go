package operations

import (
	"encoding/json"
	"io/ioutil"
)

//MenuConfig is the base of a json menu
type MenuConfig struct {
	Content Menu `json:"menu"`
}

//Menu represents the menu item
type Menu struct {
	Header string        `json:"header"`
	Items  []interface{} `json:"items"`
}

func readFile(filename string) ([]byte, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return content, nil
}

//ParseSampleJSON parses the sample json storred in assets and returns a MenuConfig struct
func ParseSampleJSON() (*MenuConfig, error) {
	var data MenuConfig

	content, err := readFile("assets/sample.json")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(content, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
