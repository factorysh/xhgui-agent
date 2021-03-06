package agent

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/davecgh/go-spew/spew"

	"github.com/factorysh/xhgui-agent/fixedqueue"
	"github.com/factorysh/xhgui-agent/metrics"
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
		log.WithField("url", mongoURL).WithError(err).Error()
		return nil, err
	}
	log.WithFields(log.Fields{
		"mongodb_host":     p.Hostname(),
		"mongodb_user":     p.User,
		"mongodb_database": p.Path,
		"queue_size":       queueSize,
	}).Info("Start agent")
	agent := &Agent{
		queue:    fixedqueue.New(queueSize),
		mongoURL: mongoURL,
		database: p.Path[1:],
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
			doc_ := agent.queue.BlPop()
			doc, ok := doc_.(*map[string]interface{})
			if !ok {
				spew.Dump(doc_)
				continue
			}
			(*doc)["_id"] = bson.NewObjectId()
			l := log.WithField("document", doc)
			err = agent.mongodb.DB(agent.database).C("results").Insert(doc)
			if err != nil {
				l.WithError(err).Error("Insert error")
			} else {
				l.Info("Inserting document")
			}
		}
	}()
	return agent, nil
}

// Handle http endpoint
func (a *Agent) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("content-type", "text/plain")
		fmt.Fprint(w, `

      _                 _                              _
__  _| |__   __ _ _   _(_)       __ _  __ _  ___ _ __ | |_
\ \/ / '_ \ / _' | | | | |_____ / _' |/ _' |/ _ \ '_ \| __|
 >  <| | | | (_| | |_| | |_____| (_| | (_| |  __/ | | | |_
/_/\_\_| |_|\__, |\__,_|_|      \__,_|\__, |\___|_| |_|\__|
            |___/                     |___/

			`)
		return
	}
	w.Header().Set("content-type", "application/json")
	if r.Method != "POST" {
		w.WriteHeader(405)
		fmt.Fprintf(w, `{"error":"Bad method: %s"}`, r.Method)
		return
	} else {
		// Increment AgentConnexion metric
		metrics.AgentConnexion.Inc()
	}
	body, err := ioutil.ReadAll(r.Body)
	// Add BodySize metric
	metrics.BodySize.Observe(float64(len(body)))
	if err != nil {
		log.Error(err)
		w.WriteHeader(500)
		fmt.Fprint(w, `{"error":"Can't read body"}`)
	}
	var msg map[string]interface{}

	err = json.Unmarshal(body, &msg)
	if err != nil {
		log.Error(err)
		w.WriteHeader(500)
		fmt.Fprint(w, `{"error":"Can't parse JSON"}`)
	}
	a.queue.Push(&msg)
	log.Info("Http handles one document")
	fmt.Fprint(w, "{}")
}
