package main

import "path/filepath"
import "io/ioutil"
import "strconv"
import "os"

func toInt(i interface{}) int {
	if i == nil {
		return 0
	}

	s := i.(string)
	num, e := strconv.Atoi(s)

	if e != nil {
		return 0
	}

	return num
}

func toBool(i interface{}) bool {
	if i == nil {
		return false
	}

	return i.(bool)
}

func toString(i interface{}) string {
	if i == nil {
		return ""
	}

	return i.(string)
}

func readfile(path string) []byte {
	if !filepath.IsAbs(path) {
		cwd, _ := os.Getwd()
		path = filepath.Join(cwd, path)
	}
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return bytes
}
