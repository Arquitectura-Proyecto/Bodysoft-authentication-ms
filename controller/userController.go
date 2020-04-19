package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jpbmdev/Bodysoft-authentication-ms/models"
	"github.com/jpbmdev/Bodysoft-authentication-ms/services"
)

func createUserController(w http.ResponseWriter, r *http.Request) {
	var User models.User
	reqBody, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(reqBody, &User); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if value := services.ValidateData(User); value != "ok" {
		http.Error(w, value, http.StatusBadRequest)
		return
	}
	if status, err := services.FindUserByEmail(User); err != nil {
		http.Error(w, err.Error(), status)
		return
	}
	if status, err := services.CreateUserAndVerificationEmail(User); err != nil {
		http.Error(w, err.Error(), status)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func assignProfileController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	token := vars["token"]
	var AuthToken models.AuthToken
	newtoken, status, err := services.AssignProfile(token)
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	}
	AuthToken.Token = newtoken
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(AuthToken)
}

func recoverPasswordController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email := vars["email"]
	if !services.ValidateEmail(email) {
		http.Error(w, "Email no valido", http.StatusBadRequest)
		return
	}
	if status, err := services.GenerateEmailData(email); err != nil {
		http.Error(w, err.Error(), status)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func verifyAcountController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email := vars["email"]
	vcode, err := strconv.Atoi(vars["vcode"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if status, err := services.ValidateUser(email, uint(vcode)); err != nil {
		http.Error(w, err.Error(), status)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func validateAuthTokenController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	token := vars["token"]
	id, typeid, profile, status, err := services.ValidateJWT(token)
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	}
	sessionData := models.SessionData{ID: id, TypeID: typeid, Profile: profile}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sessionData)
}

func authenticationController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email := vars["email"]
	password := vars["password"]
	var AuthToken models.AuthToken
	if !services.ValidateEmail(email) {
		http.Error(w, "Email no valido", http.StatusBadRequest)
		return
	}
	id, typeid, profile, status, err := services.GetAuthTockenData(email, password)
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	}
	token, err := services.GenerateJWT(id, typeid, profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	AuthToken.Token = token
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(AuthToken)
}

func chagePasswordController(w http.ResponseWriter, r *http.Request) {
	var ChangePass models.ChangePass
	reqBody, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(reqBody, &ChangePass); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err, status := services.ValidateNewPass(ChangePass.NewPassword, ChangePass.Password); err != "ok" {
		http.Error(w, err, status)
		return
	}
	if status, err := services.UpdatePassword(ChangePass); err != nil {
		http.Error(w, err.Error(), status)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
