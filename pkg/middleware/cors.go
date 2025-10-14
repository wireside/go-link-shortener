package middleware

import (
	"net/http"
	"strings"
)

func CORS(next http.Handler, allowedOrigins string, allowCredentials bool) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			origin := r.Header.Get("Origin")

			if origin == "" {
				next.ServeHTTP(w, r)
				return
			}

			allowedAll := false

			trimmed := strings.TrimSpace(allowedOrigins)
			if len(trimmed) == 0 {
				allowedAll = true
			} else {
				for _, o := range strings.Split(trimmed, ",") {
					if strings.TrimSpace(o) == "*" {
						allowedAll = true
						break
					}
				}
			}

			header := w.Header()

			if allowedAll {
				if allowCredentials {
					header.Set("Access-Control-Allow-Origin", origin)
					header.Add("Vary", "Origin")
					header.Set("Access-Control-Allow-Credentials", "true")
				} else {
					header.Set("Access-Control-Allow-Origin", "*")
				}
			} else {
				allowed := false
				for _, o := range strings.Split(allowedOrigins, ",") {
					if strings.TrimSpace(o) == origin {
						allowed = true
					}
				}

				if allowed {
					header.Set("Access-Control-Allow-Origin", origin)
					header.Add("Vary", "Origin")
					if allowCredentials {
						header.Set("Access-Control-Allow-Credentials", "true")
					}
				}
			}

			if r.Method == http.MethodOptions {
				header.Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,HEAD")
				header.Set("Access-Control-Max-Age", "86400")

				if req := r.Header.Get("Access-Control-Request-Headers"); req != "" {
					header.Set("Access-Control-Allow-Headers", req)
					header.Add("Vary", "Access-Control-Request-Headers")
				} else {
					header.Set(
						"Access-Control-Allow-Headers",
						"Authorization,Content-Type,Content-Length",
					)
				}

				w.WriteHeader(http.StatusNoContent)
				return
			}

			next.ServeHTTP(w, r)
		},
	)
}
