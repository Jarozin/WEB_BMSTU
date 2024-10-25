package middleware

import (
	"net/http"
	"project/internal/pkg/metrics"
	"strconv"
	"time"

	"github.com/urfave/negroni"
)

func PromMetrics(h http.Handler, mt metrics.PrometheusMetrics) http.Handler {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rWriterWithCode := negroni.NewResponseWriter(w)

		start := time.Now()
		h.ServeHTTP(rWriterWithCode, r)

		code := rWriterWithCode.Status()

		mt.ExecutionTime().WithLabelValues(strconv.Itoa(code), r.URL.String(), r.Method).Observe(time.Since(start).Seconds())

		mt.ExecutionTimeHist().
			WithLabelValues(strconv.Itoa(code), r.URL.String(), r.Method).
			Observe(float64(time.Since(start).Milliseconds()))

		mt.TotalHits().Inc()

		if 200 <= code && code <= 399 {
			mt.SuccessHits().
				WithLabelValues(strconv.Itoa(code), r.URL.String(), r.Method).Inc()
		} else {
			mt.ErrorsHits().
				WithLabelValues(strconv.Itoa(code), r.URL.String(), r.Method).Inc()
		}
	})
	return handler
}
