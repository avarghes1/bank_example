package controllers

import (
	"encoding/json"
	"net/http"

	"bank-example/internal/models"
	"bank-example/internal/utils"
)

// create user
var CreateUser = func(w http.ResponseWriter, r *http.Request) {
	user := &models.Users{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request"))
		return
	}

	resp := user.Create()
	utils.Respond(w, resp)
}

// update user
var UpdateUser = func(w http.ResponseWriter, r *http.Request) {
	user := &models.Users{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request"))
		return
	}

	auth := r.Context().Value("auth").(*models.Auth)
	resp := user.Update(auth.ID)
	utils.Respond(w, resp)
}

// get user
var GetUser = func(w http.ResponseWriter, r *http.Request) {
	user := &models.Users{}
	auth := r.Context().Value("auth").(*models.Auth)
	resp := user.Get(auth.ID)
	utils.Respond(w, resp)
}

// get user's balance
var GetBalance = func(w http.ResponseWriter, r *http.Request) {
	auth := r.Context().Value("auth").(*models.Auth)
	user := &models.Users{}

	resp := user.GetBalance(auth.ID)
	utils.Respond(w, resp)
}

// authorize users transaction
var AuthorizeTransaction = func(w http.ResponseWriter, r *http.Request) {
	transaction := &models.Transaction{}
	err := json.NewDecoder(r.Body).Decode(transaction)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error while decoding request body"))
		return
	}

	user := &models.Users{}
	auth := r.Context().Value("auth").(*models.Auth)
	user.ID = auth.ID
	resp := user.AuthorizeTransaction(transaction)
	utils.Respond(w, resp)
}
