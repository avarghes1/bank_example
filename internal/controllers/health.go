package controllers

import (
	"net/http"

	"bank-example/internal/utils"
)

// health check
var Health = func(w http.ResponseWriter, r *http.Request) {
	utils.Respond(w, map[string]interface{}{"health": "ok"})
}
