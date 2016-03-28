package main

import "testing"

func TestMatch(t *testing.T) {
	equal(match("api.github.com:3000", "api.github.com:3000"), true)
	equal(match("api.github.com:80", "api.github.com:3000"), false)
	equal(match("api.github.com:80", "api.github.com"), false)
	equal(match("*.github.com", "api.github.com"), true)
	equal(match("api.github.com", "github.com"), false)
	equal(match("github.com", "github.com"), true)
	equal(match("/api/*", "/api/v1/status"), true)
	equal(match("/api/*", "/api/status"), true)
	equal(match("/api/*", "/status"), false)
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
}

func TestRewriteNamedParams(t *testing.T) {
	route := newRoute("/from/:one/to/:two")

	fixtures := []struct {
		path    string
		match   bool
		pattern string
		expect  string
	}{
		{"/from/to", false, "", ""},
		{"/from/a/to/b", true, "/one/two", "/one/two"},
		{"/from/a/to/b", true, "/from/one/to/two", "/from/one/to/two"},
		{"/from/1/to/2", true, "/one/:one/two/:two", "/one/1/two/2"},
		{"/from/a/to/b", true, "/:one/:two", "/a/b"},
	}

	for _, f := range fixtures {
		match, _ := route.Match(f.path)
		equal(match, f.match)
		if match {
			equal(route.RewriteNamedParams(f.path, f.pattern), f.expect)
		}
	}
}
