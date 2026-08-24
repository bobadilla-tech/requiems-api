package middleware

import "net/http"

// CORS is minimal, wildcard-origin CORS for the public /v1/* surface.
// requiemsapi.com and requiems.xyz are different root domains (not
// subdomains of one site), so client-side JS on the dashboard calling the
// API directly is a real cross-origin request. Wildcard matches the
// retired Cloudflare Worker's prior behavior and keeps arbitrary
// third-party browser callers working, which the public docs already imply
// is supported.
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "requiems-api-key, Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
