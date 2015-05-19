package main

import "testing"

func TestParseYaml(t *testing.T) {
	file := "./fixture/config.yml"
	parseYaml(file)
}
