#!/usr/bin/env bash

usage() {
  echo "usage: {fmt|test}"
}

fmt() {
  go fmt ./...
}

test() {
  fmt
  go test ./...
}

case "$1" in
  fmt)
    fmt
  ;;
  test)
    test
  ;;
  *)
    usage
  ;;
esac
