package main

import "reflect"
import "fmt"

// utils for test

func equal(args ...interface{}) {
	if len(args)%2 != 0 {
		panic(fmt.Errorf("not matched args"))
	}

	step := len(args) / 2
	for i := 0; i < step; i++ {
		if !reflect.DeepEqual(args[i], args[i+step]) {
			panic(fmt.Errorf("expect %v to equal %v", args[i], args[i+step]))
		}
	}
}
