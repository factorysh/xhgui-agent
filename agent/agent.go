package agent

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/factorysh/xhgui-agent/fixedqueue"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	log "github.com/sirupsen/logrus"
)

//Agent for xhgui
type Agent struct {
	queue    *fixedqueue.Queue
	mongoURL string
	database string
	mongodb  *mgo.Session
}

// New Agent
func New(ctx context.Context, queueSize int, mongoURL string) (*Agent, error) {
	p, err := url.Parse(mongoURL)
	if err != nil {
		return nil, err
	}
	log.WithFields(log.Fields{
		"mongodb_host":     p.Hostname,
		"mongodb_user":     p.User,
		"mongodb_database": p.Path,
		"queue_size":       queueSize,
	}).Info("Start agent")
	agent := &Agent{
		queue:    fixedqueue.New(queueSize),
		mongoURL: mongoURL,
		database: p.Path,
	}
	go func() {
		var err error
		for {
			if agent.mongodb == nil {
				agent.mongodb, err = mgo.Dial(mongoURL)
				if err != nil {
					log.WithError(err).WithField("url", mongoURL).Error("Can't dial mongo")
					time.Sleep(10 * time.Second)
					continue
				}
			}
			doc := agent.queue.BlPop()
			err = agent.mongodb.DB(agent.database).C("plop").Insert(doc)
			if err != nil {
				log.WithError(err).WithField("document", doc).Error("Insert error")
			}
		}
	}()
	return agent, nil
}

// JSON2Bson convert a JSON document to a BSON document
func JSON2Bson(in []byte) (out []byte, err error) {
	var trace interface{}
	err = bson.UnmarshalJSON(in, &trace)
	if err != nil {
		return nil, err
	}
	return bson.Marshal(trace)
}

// Handle http endpoint
func (a *Agent) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Write([]byte(`

      _                 _                              _
__  _| |__   __ _ _   _(_)       __ _  __ _  ___ _ __ | |_
\ \/ / '_ \ / _' | | | | |_____ / _' |/ _' |/ _ \ '_ \| __|
 >  <| | | | (_| | |_| | |_____| (_| | (_| |  __/ | | | |_
/_/\_\_| |_|\__, |\__,_|_|      \__,_|\__, |\___|_| |_|\__|
            |___/                     |___/

			`))
		return
	}
	if r.Method != "POST" {
		w.WriteHeader(400)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(err)
		w.WriteHeader(500)
	}
	out, err := JSON2Bson(body)
	if err != nil {
		log.Error(err)
		w.WriteHeader(500)
	}
	a.queue.Push(out)
	w.Write(out)
}
