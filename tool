#!/usr/bin/env bash

usage() {
  echo "usage: {fmt|test|run}"
}

fmt() {
  go fmt ./...
}

run() {
  fmt
  go run main.go proxy.go util.go
}

test() {
  fmt
  go test ./...
}

case "$1" in
  fmt)
    fmt
  ;;
  run)
    run
  ;;
  test)
    test
  ;;
  *)
    usage
  ;;
esac
