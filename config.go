package main

type Rule struct {
	Host    string // host or host:port
	Path    string
	To      map[string]string
	Headers map[string]string
}

type Config struct {
	Rules []Rule
}
