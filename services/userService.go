package services

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"regexp"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jpbmdev/Bodysoft-authentication-ms/credentials"
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

func generateVcode() uint {
	v := rand.Intn(9999-1000) + 1000
	u := uint(v)
	return u
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

// FindUserByEmail ..
func FindUserByEmail(user models.User) (int, error) {
	if err := repository.FindUserByEmail(user); err == nil {
		return http.StatusConflict, errors.New("Otra cuenta ya utiliza ese correo")
	} else if err.Error() == "record not found" {
		return http.StatusCreated, nil
	} else {
		return http.StatusInternalServerError, err
	}
}

// GetAuthTockenData ..
func GetAuthTockenData(email string, pass string) (uint, uint, int, error) {
	var user models.User
	user.Email = email
	user.Password = pass
	if err := repository.GetUserByEmailPass(&user); err != nil {
		if err.Error() == "record not found" {
			return 0, 0, http.StatusUnauthorized, errors.New("Usuario o Contraseña incorrectos")
		}
		return 0, 0, http.StatusInternalServerError, err
	}
	if !user.Check {
		return 0, 0, http.StatusNotAcceptable, errors.New("La cuenta no ha sido verificada")
	}
	id := user.ID
	typeid := user.TypeID
	return id, typeid, http.StatusOK, nil
}

// CreateUserAndVerificationEmail ..
func CreateUserAndVerificationEmail(user models.User) (int, error) {
	user.Check = false
	user.Profile = false
	user.VCode = generateVcode()
	if err := repository.CreateUser(user); err != nil {
		return http.StatusInternalServerError, err
	}
	if err := repository.GenerateValidationEmail(user.Email, user.VCode); err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusCreated, nil
}

// UpdatePassword ..
func UpdatePassword(changePass models.ChangePass) (int, error) {
	var user models.User
	id, err := getIDfromJWT(changePass.Token)
	if err != nil {
		return http.StatusUnauthorized, err
	}
	user.ID = id
	user.Password = changePass.Password
	if err := repository.GetUserByIDPass(&user); err != nil {
		if err.Error() == "record not found" {
			return http.StatusUnauthorized, errors.New("Contraseña Incorrecta")
		}
		return http.StatusInternalServerError, err
	}
	user.Password = changePass.NewPassword
	if err := repository.UpdateUser(user); err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusNoContent, nil
}

// ValidateUser ..
func ValidateUser(email string, vcode uint) (int, error) {
	var user models.User
	user.Email = email
	if err := repository.GetUserByEmail(&user); err != nil {
		if err.Error() == "record not found" {
			return http.StatusConflict, err
		}
		return http.StatusInternalServerError, err
	}
	if !user.Check {
		if user.VCode == vcode {
			user.Check = true
			if err := repository.UpdateUser(user); err != nil {
				return http.StatusInternalServerError, err
			}
			return http.StatusNoContent, nil
		}
		return http.StatusUnauthorized, errors.New("Código de Verificación Incorrecto")
	}
	return http.StatusConflict, errors.New("El usuario ya esta Verificado")
}

// GenerateEmailData ..
func GenerateEmailData(email string) (int, error) {
	var user models.User
	user.Email = email
	if err := repository.GetUserByEmail(&user); err != nil {
		if err.Error() == "record not found" {
			return http.StatusConflict, err
		}
		return http.StatusInternalServerError, err
	}
	if err := repository.GenerateRecoveryEmail(email, user.Password); err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusNoContent, nil
}

// GenerateJWT ..
func GenerateJWT(id uint, typeid uint) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["ID"] = id
	claims["TypeID"] = typeid
	claims["exp"] = time.Now().Add(time.Minute * 120).Unix()
	tokenString, err := token.SignedString(credentials.JWTkey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ValidateJWT ..
func ValidateJWT(Token string) (uint, uint, int, error) {
	token, err := jwt.Parse(Token, func(tocker *jwt.Token) (interface{}, error) {
		if _, ok := tocker.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return credentials.JWTkey, nil
	})
	if err != nil {
		return 0, 0, http.StatusUnauthorized, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return uint(claims["ID"].(float64)), uint(claims["TypeID"].(float64)), http.StatusOK, nil
	}
	return 0, 0, http.StatusInternalServerError, err
}

func getIDfromJWT(token string) (uint, error) {
	id, _, _, err := ValidateJWT(token)
	if err != nil {
		return 0, err
	}
	return id, nil
}
