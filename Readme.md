
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
```

### os config

* set http proxy on `osx`

```bash
# set proxy
networksetup -setwebproxy Wi-Fi your_host your_port

# get status
networksetup -getwebproxy Wi-Fi

# close proxy
networksetup -setwebproxystate Wi-Fi off

# open proxy
networksetup -setwebproxystate Wi-Fi on
```

* `linux`

```bash
```

### License
MIT
