package main

import "testing"

func TestLogf(t *testing.T) {
	logf("name: %s", "hproxy")
	logf("version: %d", 1)
}

func TestLogBody(t *testing.T) {
	logBody("application/json; charset=utf-8", []byte(`{"name":"hproxy"}`))
}
