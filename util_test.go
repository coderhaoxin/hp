package main

import "strings"
import "testing"

func TestConvert(t *testing.T) {
	equal(toInt(nil), 0)
	equal(toBool(nil), false)
	equal(toString(nil), "")
}

func TestReadfile(t *testing.T) {
	bytes := readfile("tool")
	data := string(bytes)

	if strings.Split(data, "\n")[0] != "#!/usr/bin/env bash" {
		t.Fail()
	}
}
