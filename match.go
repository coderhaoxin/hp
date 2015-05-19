package main

import "net/url"
import "strings"
import "regexp"
import "fmt"

func urlMatch(rule Rule, uri *url.URL) bool {
	if !match(rule.Host, uri.Host) {
		return false
	}

	if strings.Contains(rule.Path, "*") {
		if match(rule.Path, uri.Path) {
			return true
		}
	}
	if strings.Contains(rule.Path, ":") {
		r := newRoute(rule.Path)
		m, _ := r.Match(uri.Path)

		if m {
			return true
		}
	}

	if rule.Path == uri.Path {
		return true
	}

	return false
}

func match(pattern, host string) bool {
	pattern = regexp.QuoteMeta(pattern)
	pattern = strings.Replace(pattern, "\\*", ".*?", -1)
	pattern = strings.Replace(pattern, ",", "|", -1)
	pattern = "^(" + pattern + ")$"
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(host)
}

type route struct {
	reg     *regexp.Regexp
	pattern string
	name    string
}

var reg1 = regexp.MustCompile(`:[^/#?()\.\\]+`)
var reg2 = regexp.MustCompile(`\*\*`)

func newRoute(pattern string) *route {
	r := route{nil, pattern, ""}
	pattern = reg1.ReplaceAllStringFunc(pattern, func(m string) string {
		return fmt.Sprintf(`(?P<%s>[^/#?]+)`, m[1:])
	})
	var index int
	pattern = reg2.ReplaceAllStringFunc(pattern, func(m string) string {
		index++
		return fmt.Sprintf(`(?P<_%d>[^#?]*)`, index)
	})
	pattern += `\/?`
	r.reg = regexp.MustCompile(pattern)
	return &r
}

func (r route) Match(path string) (bool, map[string]string) {
	matches := r.reg.FindStringSubmatch(path)

	if len(matches) > 0 && matches[0] == path {
		params := make(map[string]string)
		for i, name := range r.reg.SubexpNames() {
			if len(name) > 0 {
				params[name] = matches[i]
			}
		}
		return true, params
	}

	return false, nil
}

var urlreg = regexp.MustCompile(`:[^/#?()\.\\]+|\(\?P<[a-zA-Z0-9]+>.*\)`)

func (r *route) URLWith(args ...string) string {
	if len(args) > 0 {
		count := len(args)
		i := 0
		url := urlreg.ReplaceAllStringFunc(r.pattern, func(m string) string {
			var v interface{}
			if i < count {
				v = args[i]
			} else {
				v = m
			}
			i += 1
			return fmt.Sprintf(`%v`, v)
		})

		return url
	}

	return r.pattern
}
