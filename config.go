package main

import "github.com/pkg4go/execx"
import "runtime"

// os proxy config

func myos() string {
	return runtime.GOOS
}

func setProxy(proxyType, proxyHost, proxyPort, proxyState string) {
	if proxyType != "" || proxyHost != "" || proxyPort != "" {
		// set proxy
		if proxyType == "" {
			proxyType = "Wi-Fi"
		}
		if proxyHost == "" {
			proxyHost = "localhost"
		}
		if proxyPort == "" {
			proxyPort = "10086"
		}

		if myos() == "darwin" {
			str, err := execx.Run("networksetup", "-setwebproxy", proxyType, proxyHost, proxyPort)
			if err != nil {
				panic(err)
			}
			logf("set proxy result: %s", str)
		}
	}

	if proxyState != "" {
		setProxyState(proxyType, proxyState)
	}
}

func getProxyStatus(proxyType string) {
	if proxyType == "" {
		proxyType = "Wi-Fi"
	}

	if myos() == "darwin" {
		str, err := execx.Run("networksetup", "-getwebproxy", proxyType)
		if err != nil {
			panic(err)
		}
		logf("%s", str)
	}
}

func setProxyState(proxyType, enable string) {
	if proxyType == "" {
		proxyType = "Wi-Fi"
	}

	if myos() == "darwin" {
		str, err := execx.Run("networksetup", "-setwebproxystate", proxyType, enable)
		if err != nil {
			panic(err)
		}
		logf("%s", str)
	}
}
