package metrics

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

var (
	AgentConnexion = promauto.NewCounter(prometheus.CounterOpts{
		Name: "agent_connexion",
		Help: "Number of agent connexions",
	})
	BodySize = promauto.NewHistogram(prometheus.HistogramOpts{
		Name: "body_size",
		Help: "The size of the received body",
	})
)

func ListenAndServe(listen string) error {
	http.Handle("/metrics", promhttp.Handler())
	log.WithFields(
		log.Fields{
			"url": fmt.Sprintf("http://%s/metrics", listen),
		},
	).Info("metrics.ListenAndServe")
	return http.ListenAndServe(listen, nil)
}
