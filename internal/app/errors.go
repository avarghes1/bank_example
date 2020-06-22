package app

import (
	"net/http"

	"bank-example/internal/utils"
)

var NotFoundHandler = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		utils.Respond(w, utils.Message(false, "This resources was not found on our server"))
		next.ServeHTTP(w, r)
	})
}
