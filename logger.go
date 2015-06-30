package main

import "encoding/json"
import "io/ioutil"
import "net/http"
import "strings"
import "bytes"
import "mime"
import "fmt"
import "io"

var colors = map[string]string{
	"yellow": "\033[33m",
	"green":  "\033[32m",
	"pink":   "\033[35m",
	"red":    "\033[31m",
	"none":   "",
}
var end = "\033[0m\n"

var reqStore = make(map[int64]*http.Request)

type logOpts struct {
	sid    int64
	filter string
}

func logReq(opts logOpts, req *http.Request) {
	if opts.filter != "" && !match(opts.filter, req.URL.String()) {
		return
	}

	reqStore[opts.sid] = req
}

func logRes(opts logOpts, res *http.Response) {
	// log req

	req, ok := reqStore[opts.sid]
	if !ok {
		return
	}

	uri := req.URL.String()
	logf("red", "\nRequest - %s %s", req.Method, uri)

	logHeader(req.Header)

	method := req.Method
	if method != "GET" && method != "HEAD" {
		body, err := ioutil.ReadAll(req.Body)
		req.Body.Close()
		logErr(err)

		logBody(req.Header.Get("Content-Type"), body)
	}

	// log res

	logf("pink", "\nResponse - status: %s, proto: %s", res.Status, res.Proto)

	logHeader(res.Header)

	res.Body = NewTeeReadCloser(res.Body, writerLogger{
		data:        bytes.NewBuffer(nil),
		contentType: res.Header.Get("Content-Type"),
	})

	fmt.Print("\n\n")
}

func logHeader(header http.Header) {
	var h = map[string]string{}

	for k, v := range header {
		h[k] = strings.Join(v, " ")
	}
	data, err := json.Marshal(h)
	logErr(err)

	var out bytes.Buffer
	err = json.Indent(&out, data, "", "  ")
	logErr(err)

	logf("none", "\n%s", out.String())
}

func logBody(contentType string, body []byte) {
	if contentType == "" || len(body) == 0 {
		return
	}

	t := getType(contentType)

	switch t {
	case "json":
		var out bytes.Buffer
		err := json.Indent(&out, body, "", "  ")
		logErr(err)

		logf("none", "\n%s", out.String())
	}
}

func logErr(err error) {
	if err != nil {
		logf("yellow", "err: %v", err)
	}
}

func logf(color string, format string, args ...interface{}) {
	fmt.Printf(colors[color]+format+end, args...)
}

// writerLogger

type writerLogger struct {
	data        *bytes.Buffer
	contentType string
}

func (w writerLogger) Write(b []byte) (int, error) {
	return w.data.Write(b)
}

func (w writerLogger) Close() error {
	logBody(w.contentType, w.data.Bytes())
	return nil
}

// TeeReadCloser

type TeeReadCloser struct {
	r io.Reader
	w io.WriteCloser
	c io.Closer
}

func (t *TeeReadCloser) Read(b []byte) (int, error) {
	return t.r.Read(b)
}

func (t *TeeReadCloser) Close() error {
	e1 := t.c.Close()
	e2 := t.w.Close()

	if e1 != nil {
		return e1
	}

	return e2
}

func NewTeeReadCloser(r io.ReadCloser, w io.WriteCloser) io.ReadCloser {
	return &TeeReadCloser{io.TeeReader(r, w), w, r}
}

// getType

func getType(contentType string) string {
	mediaType, _, err := mime.ParseMediaType(contentType)
	logErr(err)

	return contentTypes[mediaType]
}

var contentTypes = map[string]string{
	"application/json-patch+json": "json",
	"application/vnd.api+json":    "json",
	"application/csp-report":      "json",
	"application/ld+json":         "json",
	"application/json":            "json",
}
