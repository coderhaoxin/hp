
[![Build status][travis-img]][travis-url]
[![License][license-img]][license-url]

### hp

A command line tool for http proxy. :dancer:

### Install

```bash
$ go get github.com/coderhaoxin/hp
```

### Dependencies

* [elazarl/goproxy](https://github.com/elazarl/goproxy)
* [haoxins/wsproxy](https://github.com/haoxins/wsproxy)
* [docopt/docopt.go](https://github.com/docopt/docopt.go)
* [go-yaml/yaml](https://github.com/go-yaml/yaml)
* [mitchellh/go-homedir](https://github.com/mitchellh/go-homedir)
* [tj/go-debug](https://github.com/tj/go-debug)

### Usage

```bash
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
```

```bash
# config file: support json or yaml

$ hp --config your-config.json

# or

$ hp --config your-config.yml
```

### Config

* yaml

```yaml
rules:
- host: localhost
  path: /api/v1/:type
  to:
    host: localhost:3001
    path: /api/:type
  # will add the headers in request
  headers:
    X-HP-A: hello
    X-HP-B: world
- host: localhost:3001
  path: /api/*
  to:
    type: origin
  headers:
    X-HP: true
- host: localhost:4000
  path: "*"
  to:
    host: httpbin.org
    path: /get
  headers:
    X-HP: true
- host: example.org
  path: "*"
  type: response
  # will add the headers in response
  headers:
    Access-Control-Allow-Origin: "*"
```

* json

```js
{
  "rules": [{
    "host": "localhost",
    "path": "/api/v1/:type",
    "to": {
      "host": "localhost:3001",
      "path": "/api/:type"
    },
    "headers": {
      "X-HP-A": "hello",
      "X-HP-B": "world"
    }
  }, {
    "host": "localhost:3001",
    "path": "/api/*",
    "to": {
      "type": "origin"
    },
    "headers": {
      "X-HP": "true"
    }
  }, {
    "host": "localhost:4000",
    "path": "*",
    "to": {
      "host": "httpbin.org",
      "path": "/get"
    },
    "headers": {
      "X-HP": "true"
    }
  }, {
    "host": "example.org",
    "path": "*",
    "type": "response",
    "headers": {
      "Access-Control-Allow-Origin": "*"
    }
  }]
}
```

### License
MIT

[travis-img]: https://img.shields.io/travis/coderhaoxin/hp.svg?style=flat-square
[travis-url]: https://travis-ci.org/coderhaoxin/hp
[license-img]: http://img.shields.io/badge/license-MIT-green.svg?style=flat-square
[license-url]: http://opensource.org/licenses/MIT
