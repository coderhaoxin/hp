package main

import "testing"

func TestParseJSON(t *testing.T) {
	file := "./fixture/config.json"
	config := parseJSON(file)
	t.Log("config: %v", config)
	// TODO: deep equal
	equal(len(config.Rules), 2)
}
