package main

import "net/http"
import "net/url"
import "strings"

type Rule struct {
	Host    string // host or host:port
	Path    string
	Type    string // request or response
	To      map[string]string
	Headers map[string]string
}

type Config struct {
	Rules []Rule
}

func (r Rule) urlMatch(uri *url.URL) bool {
	if !match(r.Host, uri.Host) {
		return false
	}

	if strings.Contains(r.Path, "*") {
		if match(r.Path, uri.Path) {
			return true
		}
	}

	if strings.Contains(r.Path, ":") {
		r := newRoute(r.Path)
		m, _ := r.Match(uri.Path)

		if m {
			return true
		}
	}

	if r.Path == uri.Path {
		return true
	}

	return false
}

func (r Rule) getTo() (toType, toHost, toPath string) {
	toType = r.To["type"]
	toHost = r.To["host"]
	toPath = r.To["path"]

	if toType == "origin" {
		toHost = ""
		toPath = ""
	} else if toHost == "" && toPath == "" {
		toType = "origin"
	}

	return
}

func (r Rule) setHeaders(req *http.Request) {
	for name, value := range r.Headers {
		debug("set header - %s : %s", name, value)
		req.Header.Set(name, value)
	}
}

func (r Rule) setResHeaders(res *http.Response) {
	for name, value := range r.Headers {
		debug("set header - %s : %s", name, value)
		res.Header.Set(name, value)
	}
}

func (r Rule) redirect(req *http.Request) {
	toType, toHost, toPath := r.getTo()
	debug("toType: %s, toHost: %s, toPath: %s", toType, toHost, toPath)

	if toType == "origin" {
		return
	}

	if toHost != "" {
		req.URL.Host = toHost
	}

	if toPath != "" {
		if strings.Contains(toPath, ":") {
			req.URL.Path = newRoute(r.Path).RewriteNamedParams(req.URL.Path, toPath)
		} else {
			req.URL.Path = toPath
		}
	}
}
