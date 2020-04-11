package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jpbmdev/Bodysoft-authentication-ms/models"
	"github.com/jpbmdev/Bodysoft-authentication-ms/services"
)

func createUserController(w http.ResponseWriter, r *http.Request) {
	var User models.User
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &User)
	if value := services.ValidateData(User); value != "ok" {
		http.Error(w, value, http.StatusBadRequest)
		return
	}
	if err, status := services.ServFindUserByEmail(User); err != "Correo diponible" {
		http.Error(w, err, status)
		return
	}
	if value := services.ServCreateUser(User); value != "ok" {
		http.Error(w, value, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func authenticationController(w http.ResponseWriter, r *http.Request) {
	var Auth models.Auth
	vars := mux.Vars(r)
	email := vars["email"]
	password := vars["password"]
	Auth.Email = email
	Auth.Password = password
	if value := services.ValidateAuth(Auth); value != "ok" {
		http.Error(w, value, http.StatusBadRequest)
		return
	}
	if err, status := services.AuthFindUserByEmail(Auth); err != "ok" {
		http.Error(w, err, status)
		return
	}
	if err, status := services.AuthFindUserByEmailPass(Auth); err != "ok" {
		http.Error(w, err, status)
		return
	}
	tokendata, err := services.CreateAuthToker(Auth)
	if err != "ok" {
		http.Error(w, err, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tokendata)
}

func chagePasswordController(w http.ResponseWriter, r *http.Request) {
	var ChangePass models.ChangePass
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &ChangePass)
	if err, status := services.AuthFindUserByIDPass(ChangePass); err != "ok" {
		http.Error(w, err, status)
		return
	}
	if err, status := services.ValidateNewPass(ChangePass.NewPassword, ChangePass.Password); err != "ok" {
		http.Error(w, err, status)
		return
	}
	if err, status := services.UpdatePassword(ChangePass); err != "ok" {
		http.Error(w, err, status)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func recoverPasswordWithEmail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email := vars["email"]
	var Auth models.Auth
	Auth.Email = email
	if !services.ValidateEmail(email) {
		http.Error(w, "Email no valido", http.StatusBadRequest)
		return
	}
	if err, status := services.AuthFindUserByEmail(Auth); err != "ok" {
		http.Error(w, err, status)
		return
	}
	if err := services.GenerateEmailData(email); err != "ok" {
		http.Error(w, err, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
