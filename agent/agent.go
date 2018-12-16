package agent

import (
	"io/ioutil"
	"net/http"

	"github.com/factorysh/xhgui-agent/fixedqueue"
	"github.com/globalsign/mgo/bson"
	log "github.com/sirupsen/logrus"
)

type Agent struct {
	queue *fixedqueue.Queue
}

func New(queueSize int) *Agent {
	return &Agent{
		queue: fixedqueue.New(queueSize),
	}
}

// Json2Bson convert a JSON document to a BSON document
func Json2Bson(in []byte) (out []byte, err error) {
	var trace interface{}
	err = bson.UnmarshalJSON(in, &trace)
	if err != nil {
		return nil, err
	}
	return bson.Marshal(trace)
}

// Handle http endpoint
func (a *Agent) Handle(w http.ResponseWriter, r *http.Request) {
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
	a.queue.Push(out)
	w.Write(out)
}
