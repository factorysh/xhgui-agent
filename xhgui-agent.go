package main

import (
	"context"
	"net/http"
	"os"

	"github.com/factorysh/xhgui-agent/agent"
	"github.com/factorysh/xhgui-agent/metrics"
	"github.com/onrik/logrus/filename"
	log "github.com/sirupsen/logrus"
)

func main() {
	filenameHook := filename.NewHook()
	log.AddHook(filenameHook)
	listen := os.Getenv("LISTEN")
	if listen == "" {
		listen = "127.0.0.1:8002"
	}
	mongo := os.Getenv("MONGODB_URL")
	if mongo == "" {
		mongo = "mongodb://mongo:27072/xhprof"
	}
	ctx := context.Background()
	a, err := agent.New(ctx, 100, mongo)
	if err != nil {
		log.WithField("mongo", mongo).WithError(err).Fatal("Agent error")
		return
	}
	http.HandleFunc("/", a.Handle)

	go log.Fatal(metrics.ListenAndServe(listen))
	log.Fatal(http.ListenAndServe(listen, nil))
}
