package test

import (
	"fmt"
	"io"
	"net/http"

	"code.cloudfoundry.org/routing-release/gorouter/route"
	"code.cloudfoundry.org/routing-release/gorouter/test/common"
	"github.com/nats-io/nats.go"
)

func NewGreetApp(urls []route.Uri, rPort uint16, mbusClient *nats.Conn, tags map[string]string) *common.TestApp {
	app := common.NewTestApp(urls, rPort, mbusClient, tags, "")
	app.AddHandler("/", greetHandler)
	app.AddHandler("/forwardedprotoheader", headerHandler)

	return app
}

func headerHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, fmt.Sprintf("%+v", r.Header.Get("X-Forwarded-Proto")))
}
func greetHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, world")
}
