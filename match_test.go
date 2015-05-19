package main

import "testing"

func TestMatch(t *testing.T) {
	equal(match("*.github.com", "api.github.com"), true)
	equal(match("github.com", "github.com"), true)
	equal(match("/api/*", "/api/v1/status"), true)
	equal(match("/api/*", "/api/status"), true)
}

func TestRoute(t *testing.T) {
	routes := []struct {
		path   string
		match  bool
		params map[string]string
	}{
		{"/shops/10086/items/123", true, map[string]string{"shop": "10086", "item": "123"}},
		{"/shops/hello/items/world", true, map[string]string{"shop": "hello", "item": "world"}},
		{"shops/hello/items/world", false, map[string]string{}},
		{"/shops/10086/items/123/", true, map[string]string{"shop": "10086", "item": "123"}},
		{"/shops/10086/items/123//", false, map[string]string{}},
		{"/shops/10086//items/123/", false, map[string]string{}},
		{"/shops/10086/items/123/", true, map[string]string{"shop": "10086", "item": "123"}},
	}

	route := newRoute("/shops/:shop/items/:item")
	for _, r := range routes {
		match, params := route.Match(r.path)
		equal(match, r.match)
		if match {
			t.Log(params)
			equal(params["shop"], r.params["shop"])
			equal(params["item"], r.params["item"])
		}
	}

	r := route.URLWith("s", "i")
	equal(r, "/shops/s/items/i")
}
