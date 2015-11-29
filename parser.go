package main

import "gopkg.in/yaml.v2"
import "encoding/json"

func parseYaml(path string) Config {
	data := readfile(path)
	var c Config

	err := yaml.Unmarshal(data, &c)

	if err != nil {
		panic(err)
	}

	return c
}

func parseJSON(path string) Config {
	data := readfile(path)
	var c Config
	json.Unmarshal(data, &c)

	return c
}
