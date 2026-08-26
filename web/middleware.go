package web

import "net/http"

func WithRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Request-ID", "deterministic")
		next.ServeHTTP(w, r)
	})
}
func RequireMethod(method string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			http.Error(w, "method not allowed", 405)
			return
		}
		next.ServeHTTP(w, r)
	})
}
