package metrics

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type PrometheusMetrics interface {
	SetupMetrics() error
	ExecutionTimeHist() *prometheus.HistogramVec
	ErrorsHits() *prometheus.CounterVec
	SuccessHits() *prometheus.CounterVec
	TotalHits() prometheus.Counter
	ExecutionTime() *prometheus.SummaryVec
}

type prometheusMetrics struct {
	executionTime     *prometheus.SummaryVec
	executionTimeHist *prometheus.HistogramVec
	errorsHits        *prometheus.CounterVec
	successHits       *prometheus.CounterVec
	totalHits         prometheus.Counter
}

func NewPrometheusMetrics(serviceName string) PrometheusMetrics {
	metrics := &prometheusMetrics{
		executionTime: prometheus.NewSummaryVec(prometheus.SummaryOpts{
			Name: serviceName + "_durations",
			Help: "Shows durations in minutes of request execution",
		}, []string{"status", "path", "method"}),
		executionTimeHist: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Name: serviceName + "_durations_hist",
			Help: "Shows durations in minutes of request execution",
		}, []string{"status", "path", "method"}),
		errorsHits: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: serviceName + "_errors_hits",
			Help: "Counts errors responses from service",
		}, []string{"status", "path", "method"}),
		successHits: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: serviceName + "_success_hits",
			Help: "Counts success responses from service",
		}, []string{"status", "path", "method"}),
		totalHits: prometheus.NewCounter(prometheus.CounterOpts{
			Name: serviceName + "_total_hits",
			Help: "Counts all responses from service",
		}),
	}

	return metrics
}

func (m *prometheusMetrics) SetupMetrics() error {
	if err := prometheus.Register(m.executionTime); err != nil {
		return err
	}

	if err := prometheus.Register(m.executionTimeHist); err != nil {
		return err
	}

	if err := prometheus.Register(m.errorsHits); err != nil {
		return err
	}

	if err := prometheus.Register(m.successHits); err != nil {
		return err
	}

	if err := prometheus.Register(m.totalHits); err != nil {
		return err
	}

	return nil
}

func (m *prometheusMetrics) ExecutionTime() *prometheus.SummaryVec {
	return m.executionTime
}

func (m *prometheusMetrics) ExecutionTimeHist() *prometheus.HistogramVec {
	return m.executionTimeHist
}

func (m *prometheusMetrics) ErrorsHits() *prometheus.CounterVec {
	return m.errorsHits
}
func (m *prometheusMetrics) SuccessHits() *prometheus.CounterVec {
	return m.successHits
}
func (m *prometheusMetrics) TotalHits() prometheus.Counter {
	return m.totalHits
}

func ServePrometheusHTTP(addr string) {
	mux := http.NewServeMux()

	mux.Handle("/metrics", promhttp.Handler())

	err := http.ListenAndServe(addr, mux)
	if err != nil {
		log.Fatal(err)
	}
}
