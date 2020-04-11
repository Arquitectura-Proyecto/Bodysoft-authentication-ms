package services

import (
	"net/http"
	"regexp"

	"github.com/jpbmdev/Bodysoft-authentication-ms/models"
	"github.com/jpbmdev/Bodysoft-authentication-ms/repository"
)

// ValidateEmail ..
func ValidateEmail(email string) bool {
	Re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-z]{2,4}$`)
	return Re.MatchString(email)
}

// ValidatePass ..
func ValidatePass(pass string) bool {
	if len(pass) <= 7 {
		return false
	}
	return true
}

// ValidateData ..
func ValidateData(user models.User) string {
	if !ValidateEmail(user.Email) {
		return "Email no valido"
	}
	if !ValidatePass(user.Password) {
		return "Contraseña muy corta, minimo 8 caracteres"
	}
	return "ok"
}

// ValidateNewPass ..
func ValidateNewPass(newpass string, pass string) (string, int) {
	if !ValidatePass(newpass) {
		return "Contraseña muy corta, minimo 8 caracteres", http.StatusBadRequest
	}
	if newpass == pass {
		return "Las contraseña nueva debe ser diferente a la actual", http.StatusBadRequest
	}
	return "ok", http.StatusNoContent
}

// ValidateAuth ..
func ValidateAuth(auth models.Auth) string {
	if !ValidateEmail(auth.Email) {
		return "Email no valido"
	}
	return "ok"
}

// ServFindUserByEmail ..
func ServFindUserByEmail(user models.User) (string, int) {
	if status := repository.FindUserByEmail(user); status == "Usuario Encontrado" {
		return "Otra cuenta ya utiliza ese correo", http.StatusConflict
	} else if status == "record not found" {
		return "Correo diponible", http.StatusCreated
	} else {
		return status, http.StatusInternalServerError
	}
}

// AuthFindUserByEmail ..
func AuthFindUserByEmail(auth models.Auth) (string, int) {
	var user models.User
	user.Email = auth.Email
	if status := repository.FindUserByEmail(user); status == "Usuario Encontrado" {
		return "ok", http.StatusOK
	} else if status == "record not found" {
		return "Usuario no encontrado", http.StatusConflict
	} else {
		return status, http.StatusInternalServerError
	}
}

// AuthFindUserByEmailPass ..
func AuthFindUserByEmailPass(auth models.Auth) (string, int) {
	var user models.User
	user.Email = auth.Email
	user.Password = auth.Password
	if status := repository.FindUserByEmailPass(user); status == "Usuario Encontrado" {
		return "ok", http.StatusOK
	} else if status == "record not found" {
		return "Contraseña incorrecta", http.StatusUnauthorized
	} else {
		return status, http.StatusInternalServerError
	}
}

// AuthFindUserByIDPass ..
func AuthFindUserByIDPass(changePass models.ChangePass) (string, int) {
	var user models.User
	user.ID = changePass.ID
	user.Password = changePass.Password
	if status := repository.FindUserByIDPass(user); status == "Usuario Encontrado" {
		return "ok", http.StatusOK
	} else if status == "record not found" {
		return "Contraseña incorrecta", http.StatusUnauthorized
	} else {
		return status, http.StatusInternalServerError
	}
}

// ServCreateUser ..
func ServCreateUser(user models.User) string {
	return repository.CreateUser(user)
}

// CreateAuthToker ..
func CreateAuthToker(Auth models.Auth) (models.TokenData, string) {
	var user models.User
	var TokenData models.TokenData
	user.Email = Auth.Email
	if err := repository.GetUserUserByEmail(&user); err != "ok" {
		return TokenData, err
	}
	TokenData.ID = user.ID
	TokenData.TypeID = user.TypeID
	return TokenData, "ok"
}

// UpdatePassword ..
func UpdatePassword(changePass models.ChangePass) (string, int) {
	var user models.User
	user.ID = changePass.ID
	if err := repository.GetUserUserByID(&user); err != "ok" {
		return err, http.StatusInternalServerError
	}
	user.Password = changePass.NewPassword
	if err := repository.UpdateUser(user); err != "ok" {
		return err, http.StatusInternalServerError
	}
	return "ok", http.StatusNoContent
}

// GenerateEmailData ..
func GenerateEmailData(email string) string {
	var user models.User
	user.Email = email
	if err := repository.GetUserUserByEmail(&user); err != "ok" {
		return err
	}
	if err := repository.GenerateEmail(email, user.Password); err != "ok" {
		return err
	}
	return "ok"
}
