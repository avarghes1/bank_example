package app

import (
	"context"
	"net/http"

	"bank-example/internal/models"
	u "bank-example/internal/utils"
)

type exempt struct {
	path   string
	method string
}

// basic auth
var Authentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		notAuth := []exempt{
			exempt{path: "/health", method: "GET"},
			exempt{path: "/api/v1/users", method: "PUT"},
		}
		requestPath := r.URL.Path
		requestMethod := r.Method

		for _, value := range notAuth {
			if value.path == requestPath && value.method == requestMethod {
				next.ServeHTTP(w, r)
				return
			}
		}
		email, password, ok := r.BasicAuth()
		if !ok {
			response := u.Message(false, "forbidden")
			w.WriteHeader(http.StatusForbidden)
			u.Respond(w, response)
			return
		}

		auth := models.AuthUser(email, password)
		if !auth.Authenticated {
			response := u.Message(false, "forbidden")
			w.WriteHeader(http.StatusForbidden)
			u.Respond(w, response)
			return
		}

		ctx := context.WithValue(r.Context(), "auth", auth)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
		return
	})
}
