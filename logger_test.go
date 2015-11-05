package main

import "testing"

func TestLogf(t *testing.T) {
	logf("red", "name: %s", "hp")
	logf("yellow", "version: %d", 1)
}

func TestLogBody(t *testing.T) {
	logBody("application/json; charset=utf-8", []byte(`{"name":"hp"}`))
}
