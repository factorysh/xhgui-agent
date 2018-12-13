package main

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/globalsign/mgo/bson"
	log "github.com/sirupsen/logrus"
)

func Json2Bson(in []byte) (out []byte, err error) {
	var trace interface{}
	err = bson.UnmarshalJSON(in, &trace)
	if err != nil {
		return nil, err
	}
	return bson.Marshal(trace)
}

func main() {
	listen := os.Getenv("LISTEN")
	if listen == "" {
		listen = "127.0.0.1:8080"
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(400)
			return
		}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Error(err)
			w.WriteHeader(500)
		}
		out, err := Json2Bson(body)
		if err != nil {
			log.Error(err)
			w.WriteHeader(500)
		}
		w.Write(out)
	})

	log.Fatal(http.ListenAndServe(listen, nil))

}
