package main

import (
	"context"
	"net/http"
	"os"

	"github.com/factorysh/xhgui-agent/agent"
	log "github.com/sirupsen/logrus"
)

func main() {
	listen := os.Getenv("LISTEN")
	if listen == "" {
		listen = "127.0.0.1:8080"
	}
	mongo := os.Getenv("MONGODB_URL")
	if mongo == "" {
		mongo = "mongodb://mongo:27072/xhgui"
	}
	ctx := context.Background()
	a, err := agent.New(ctx, 100, mongo)
	if err != nil {
		log.WithField("mongo", mongo).WithError(err).Fatal("Agent error")
		return
	}
	http.HandleFunc("/", a.Handle)

	log.Fatal(http.ListenAndServe(listen, nil))
}
