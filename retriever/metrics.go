package retriever

import (
	"context"
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/zero-gravity-labs/zerog-data-avail/common"
)

const (
	Namespace = "zgda_retriever"
)

type MetricsConfig struct {
	HTTPPort string
}

type Metrics struct {
	registry *prometheus.Registry

	NumRetrievalRequest prometheus.Counter

	httpPort string
	logger   common.Logger
}

func NewMetrics(httpPort string, logger common.Logger) *Metrics {
	reg := prometheus.NewRegistry()
	reg.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))
	reg.MustRegister(collectors.NewGoCollector())

	metrics := &Metrics{
		registry: reg,
		NumRetrievalRequest: promauto.With(reg).NewCounter(
			prometheus.CounterOpts{
				Namespace: Namespace,
				Name:      "request",
				Help:      "the number of retrieval requests",
			},
		),
		httpPort: httpPort,
		logger:   logger,
	}
	return metrics
}

// IncrementRetrievalRequestCounter increments the number of retrieval requests
func (g *Metrics) IncrementRetrievalRequestCounter() {
	// if anyone wants to add new metrics type and use "Add" for adding float,
	// please add the lock, since that ops is not atomic
	g.NumRetrievalRequest.Inc()
}

func (g *Metrics) Start(ctx context.Context) {
	g.logger.Info("Starting metrics server at ", "port", g.httpPort)
	addr := fmt.Sprintf(":%s", g.httpPort)
	go func() {
		log := g.logger
		http.Handle("/metrics", promhttp.HandlerFor(
			g.registry,
			promhttp.HandlerOpts{},
		))
		err := http.ListenAndServe(addr, nil)
		log.Error("Prometheus server failed", "err", err)
	}()
}