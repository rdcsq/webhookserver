package middleware

import (
	"encoding/json"
	"net/http"
	"strings"
	"webhookserver/utils"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := strings.Split(r.Header.Get("Authorization"), " ")
		if len(token) != 2 || token[0] != "Bearer" {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(map[string]string{"error": "Bad Request"})
			return
		}

		origin, err := utils.ValidateJwt(token[1])
		if err != nil {
			w.WriteHeader(401)
			return
		}
		r.Header.Set("X-Origin", origin)

		next.ServeHTTP(w, r)
	})
}
