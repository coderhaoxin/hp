package main

import "encoding/json"

func parseJSON(path string) Config {
	data := readfile(path)
	var c Config
	json.Unmarshal(data, &c)

	return c
}
