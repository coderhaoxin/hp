package main

import "github.com/elazarl/goproxy/transport"
import "github.com/docopt/docopt-go"
import "github.com/elazarl/goproxy"
import . "github.com/tj/go-debug"
import "net/http"
import "strconv"
import "os"

var debug = Debug("hproxy")

const version = "0.4.0"
const usage = `
	Usage:
		hproxy [-c=<config>] [-p=<port>] [-f=<filter>] [-v] [-i]
		hproxy --help
		hproxy --version

	Options:
		-c=<config> Required, config file path
		-p=<port>   Required, listening port
		-f=<filter> Filter, filter by uri
		-v          Verbose mode
		-i          Inspect
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
	inspect := toBool(args["-i"])
	filter := toString(args["-f"])

	debug("config: %s, port: %d, verbose: %v, inspect: %v, filter: %v", confPath, port, verbose, inspect, filter)

	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = false

	debug("hproxy listening on %d", port)

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
