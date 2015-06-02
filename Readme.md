### hproxy

a command line tool for http proxy, based on [elazarl/goproxy](https://github.com/elazarl/goproxy)

### Install

```bash
$ go get github.com/coderhaoxin/hproxy
```

### Usage

```bash
$ hproxy -c your-config.yml
```

### Config

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
