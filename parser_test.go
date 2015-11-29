package main

import "testing"

func TestParseYaml(t *testing.T) {
	file := "./fixture/config.yml"
	config := parseYaml(file)
	t.Logf("config: %v", config)
	checkParsedConfig(config)
}

func TestParseJSON(t *testing.T) {
	file := "./fixture/config.json"
	config := parseJSON(file)
	t.Logf("config: %v", config)
	checkParsedConfig(config)
}

func checkParsedConfig(c Config) {
	for _, rule := range c.Rules {
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
