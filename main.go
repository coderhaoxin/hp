package main

import "github.com/elazarl/goproxy/transport"
import "github.com/docopt/docopt-go"
import "github.com/elazarl/goproxy"
import . "github.com/tj/go-debug"
import "net/http"
import "strconv"
import "os"

var debug = Debug("hp")

const version = "0.5.0"
const usage = `
	Usage:
		hp [--config=<config>] [--port=<port>] [--filter=<filter>] [--verbose] [--inspect]
		hp --help
		hp --version

	Options:
		-c --config=<config> Required, config file path
		-p --port=<port>     Required, listening port
		-f --filter=<filter> Filter, filter by uri
		-v --verbose         Verbose mode
		-i --inspect         Inspect
		-h --help            Show this screen
		--version            Show version
`

func main() {
	args, _ := docopt.Parse(usage, os.Args[1:], true, version, false)

	debug("args: %v", args)

	confPath := toString(args["--config"].(string))
	if confPath == "" {
		confPath = "config.yml"
	}

	port := toInt(args["--port"].(string))
	verbose := toBool(args["--verbose"])
	inspect := toBool(args["--inspect"])
	filter := toString(args["--filter"])

	debug("config: %s, port: %d, verbose: %v, inspect: %v, filter: %v", confPath, port, verbose, inspect, filter)

	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = false

	debug("hp listening on %d", port)

	config := parseJSON(confPath)
	debug("config: %v", config)

	tr := transport.Transport{Proxy: transport.ProxyFromEnvironment}

	proxy.OnRequest().DoFunc(func(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
		debug("on request - %s", req.URL.String())

		if inspect {
			ctx.RoundTripper = goproxy.RoundTripperFunc(func(req *http.Request, ctx *goproxy.ProxyCtx) (res *http.Response, err error) {
				ctx.UserData, res, err = tr.DetailedRoundTrip(req)
				return
			})
			logReq(logOpts{
				sid:    ctx.Session,
				filter: filter,
			}, req)
		}

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

	proxy.OnResponse().DoFunc(func(res *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
		if inspect {
			logRes(logOpts{
				sid:    ctx.Session,
				filter: filter,
			}, res)
		}

		return res
	})

	err := http.ListenAndServe(":"+strconv.Itoa(port), proxy)
	if err != nil {
		panic(err)
	}
}
