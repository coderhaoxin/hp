package main

import "testing"

func TestRule(t *testing.T) {
	// rule 01
	rule01 := Rule{
		Host: "a.com",
		Path: "/a",
		To: map[string]string{
			"host": "b.com",
			"path": "/b",
		},
		Headers: map[string]string{
			"X-Proxy-Flag": "true",
		},
	}

	t.Log(rule01)

	// TODO
}

func TestSendfile(t *testing.T) {

}
