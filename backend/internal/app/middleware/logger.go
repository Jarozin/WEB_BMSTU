package middleware

import (
	"net/http"
)

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// start := time.Now()
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
		// fmt.Printf("%s -- [%s] %s, %s %s\n",
		// 	time.Now().Format(time.DateTime), r.Method, r.RemoteAddr, r.URL.Path, time.Since(start))
	})
}
