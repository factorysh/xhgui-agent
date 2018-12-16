package main

import (
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
	a := agent.New(100)
	http.HandleFunc("/", a.Handle)

	log.Fatal(http.ListenAndServe(listen, nil))

}
