package main

import "github.com/docopt/docopt-go"
import "github.com/elazarl/goproxy"
import . "github.com/tj/go-debug"
import "net/http"
import "strconv"
import "os"

var debug = Debug("hproxy")

const version = "0.1.0"
const usage = `
	Usage:
		hproxy [-c=<config>] [-p=<port>] [-v]
		hproxy --help
		hproxy --version

	Options:
		-c=<config> Required, config file path
		-p=<port>   Required, listening port
		-v         Verbose mode
		--help      Show this screen
		--version   Show version
`

func main() {
	args, _ := docopt.Parse(usage, os.Args[1:], true, version, false)

	debug("args: %v", args)

	confPath := toString(args["-c"].(string))
	if confPath == "" {
		confPath = "config.yml"
	}

	port := toInt(args["-p"].(string))
	verbose := toBool(args["-v"])

	debug("config: %s, port: %d, verbose: %v", confPath, port, verbose)

	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = verbose

	debug("hproxy listening on %d", port)

	config := parseJSON(confPath)

	proxy.OnRequest().DoFunc(
		func(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
			debug("on request - %s", req.URL.String())
			for _, rule := range config.Rules {
				debug("check rule: %v vs url: %v", rule, req.URL)
				if rule.urlMatch(req.URL) {
					debug("matched")
					rule.setHeaders(req)
					rule.redirect(req)
				}
			}
			return req, nil
		})

	err := http.ListenAndServe(":"+strconv.Itoa(port), proxy)
	if err != nil {
		panic(err)
	}
}
