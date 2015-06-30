package main

import "net/http"
import "fmt"
import "io"

func ok(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "ok")
}

func main() {
	http.HandleFunc("/", ok)

	err := http.ListenAndServe(":3003", nil)
	fmt.Print("port: 3003")

	if err != nil {
		fmt.Print(err)
	}
}
