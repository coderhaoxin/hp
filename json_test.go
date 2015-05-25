package main

import "testing"

func TestParseJSON(t *testing.T) {
	file := "./fixture/config.json"
	config := parseJSON(file)
	t.Log("config: %v", config)

	for _, rule := range config.Rules {
		switch rule.Host {
		case "localhost":
			equal(rule.Path, "/api/v1/:type")
			toType, toHost, toPath := rule.getTo()
			equal(toType, toHost, toPath, "", "localhost:3001", "/api/:type")
		case "localhost:3001":
			equal(rule.Path, "/api/*")
			toType, toHost, toPath := rule.getTo()
			equal(toType, toHost, toPath, "origin", "", "")
		case "localhost:4000":
			equal(rule.Path, "*")
			toType, toHost, toPath := rule.getTo()
			equal(toType, toHost, toPath, "", "httpbin.org", "/get")
		}
	}
}
