package main

// import "net/url"
import "strings"
import "regexp"
import "fmt"

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

func (r *route) Match(path string) (bool, map[string]string) {
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

func (r *route) RewriteNamedParams(from, to string) string {
	fromMatches := r.reg.FindStringSubmatch(from)

	if len(fromMatches) > 0 {
		for i, name := range r.reg.SubexpNames() {
			if len(name) > 0 {
				to = strings.Replace(to, ":"+name, fromMatches[i], -1)
			}
		}
	}

	return to
}
