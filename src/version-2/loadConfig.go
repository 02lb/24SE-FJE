package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type IconFamily struct {
	Intermediate string `json:"intermediate"`
	Leaf         string `json:"leaf"`
}

type Config struct {
	IconFamilies map[string]IconFamily `json:"iconFamilies"`
}

func LoadConfig(filePath string) (*Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
