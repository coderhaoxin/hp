### hproxy

a command line tool for http proxy, based on [elazarl/goproxy](https://github.com/elazarl/goproxy)

### Install

```bash
$ go get github.com/coderhaoxin/hproxy
```

### Usage

```
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
```

```bash
# config file: support json or yaml

$ hproxy -c your-config.json

# or

$ hproxy -c your-config.yml
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
