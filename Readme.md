[![Build status][travis-img]][travis-url]
[![License][license-img]][license-url]

### hp

A command line tool for http proxy. :dancer:

### Install

```bash
$ go get github.com/coderhaoxin/hp
```

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

* json

```js
{
  "rules": [{
    "host": "example.com",
    "path": "/api/*",
    "to": {
      "host": "localhost:3000"
    }
  }, {
    "host": "localhost:3003",
    "path": "/api/v1/*",
    "headers": {
      // will add the headers
      "X-Api-Version": "v1"
    },
    "to": {
      "type": "origin"
    }
  }]
}
```

* yaml

```yaml
rules:
- host: example.com
  path: "/api/*"
  to:
    host: localhost:3000
- host: localhost:3003
  path: "/api/v1/*"
  headers:
    # will add the headers
    X-Api-Version: v1
  to:
    type: origin
```

### License
MIT

[travis-img]: https://img.shields.io/travis/coderhaoxin/hp.svg?style=flat-square
[travis-url]: https://travis-ci.org/coderhaoxin/hp
[license-img]: http://img.shields.io/badge/license-MIT-green.svg?style=flat-square
[license-url]: http://opensource.org/licenses/MIT
