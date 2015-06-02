package main

import "encoding/json"
import "math/rand"
import "io/ioutil"
import "net/http"
import "strings"
import "bytes"
import "mime"
import "fmt"
import "io"

var colors = []string{"31", "32", "33", "34", "35", "36"}
var end = "\033[0m\n"

func logf(format string, args ...interface{}) {
	color := "\033[" + colors[rand.Intn(len(colors))] + "m "
	fmt.Printf(color+format+end, args...)
}

func logReq(req *http.Request) {
	logf("request method: %s, URI: %s", req.Method, req.URL.String())

	logHeader(req.Header)

	body, err := ioutil.ReadAll(req.Body)
	req.Body.Close()
	logErr(err)

	logBody(req.Header.Get("Content-Type"), body)
}

func logRes(res *http.Response) {
	logf("statusCode: %d; status: %s, proto: %s", res.StatusCode, res.Status, res.Proto)

	logHeader(res.Header)

	res.Body = NewTeeReadCloser(res.Body, writerLogger{
		data:        bytes.NewBuffer(nil),
		contentType: res.Header.Get("Content-Type"),
	})
}

func logErr(err error) {
	if err != nil {
		logf("err: %v", err)
	}
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

	logf("*** header ***\n\n%s", out.String())
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

		logf("*** body ***\n\n%s", out.String())
	}
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
