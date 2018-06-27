package main

import "github.com/elazarl/goproxy/transport"
import "github.com/docopt/docopt-go"
import "github.com/elazarl/goproxy"
import "net/http"
import "strconv"
import "path"
import "os"

const version = "0.8.0"
const usage = `
	Usage:
		hp [--config=<config>] [--port=<port>] [--filter=<filter>] [--verbose] [--inspect]
		hp [--proxy-status] [--proxy-state=<state>] [--proxy-type=<type>] [--proxy-host=<host>] [--proxy-port=<port>]
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

		--proxy-status
		--proxy-state=<state>
		--proxy-type=<type>
		--proxy-host=<host>
		--proxy-port=<port>
`

func main() {
	args, _ := docopt.Parse(usage, os.Args[1:], true, version, false)

	parseProxyArgs(args)
	if args["--config"] == nil {
		// set proxy only
		return
	}

	confPath := toString(args["--config"].(string))
	if confPath == "" {
		confPath = "config.yml"
	}

	port := toInt(args["--port"].(string))
	// verbose := toBool(args["--verbose"])
	inspect := toBool(args["--inspect"])
	filter := toString(args["--filter"])

	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = false

	var config Config
	if ext := path.Ext(confPath); ext == "json" {
		// json
		config = parseJSON(confPath)
	} else {
		// yaml
		config = parseYaml(confPath)

	}

	tr := transport.Transport{Proxy: transport.ProxyFromEnvironment}

	proxy.OnRequest().DoFunc(func(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
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
			if rule.Type == "response" {
				return req, nil
			}
			if rule.urlMatch(req.URL) {
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

		uri := res.Request.URL

		for _, rule := range config.Rules {
			if rule.Type == "response" {
				if rule.urlMatch(uri) {
					rule.setResHeaders(res)
				}
			}
		}

		return res
	})

	err := http.ListenAndServe(":"+strconv.Itoa(port), proxy)
	if err != nil {
		panic(err)
	}
}

func parseProxyArgs(args map[string]interface{}) {
	showProxyStatus := toBool(args["--proxy-status"])
	proxyState := toString(args["--proxy-state"])
	proxyType := toString(args["--proxy-type"])
	proxyHost := toString(args["--proxy-host"])
	proxyPort := toString(args["--proxy-port"])

	setProxy(proxyType, proxyHost, proxyPort, proxyState)

	if showProxyStatus {
		getProxyStatus(proxyType)
	}
}
