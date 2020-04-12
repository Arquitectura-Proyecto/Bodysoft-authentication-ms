package services

import (
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

// ServFindUserByEmail ..
func ServFindUserByEmail(user models.User) (string, int) {
	if err := repository.FindUserByEmail(user); err == nil {
		return "Otra cuenta ya utiliza ese correo", http.StatusConflict
	} else if err.Error() == "record not found" {
		return "Correo diponible", http.StatusCreated
	} else {
		return err.Error(), http.StatusInternalServerError
	}
}

// GetAuthTockenData ..
func GetAuthTockenData(email string, pass string) (uint, uint, int, error) {
	var user models.User
	user.Email = email
	user.Password = pass
	if err := repository.GetUserByEmailPass(&user); err != nil {
		if err.Error() == "record not found" {
			return 0, 0, http.StatusUnauthorized, err
		}
		return 0, 0, http.StatusInternalServerError, err
	}
	id := user.ID
	typeid := user.TypeID
	return id, typeid, http.StatusOK, nil
}

// ServCreateUser ..
func ServCreateUser(user models.User) error {
	user.Validate = false
	user.VCode = generateVcode()
	return repository.CreateUser(user)
}

// UpdatePassword ..
func UpdatePassword(changePass models.ChangePass) (int, error) {
	var user models.User
	user.ID = changePass.ID
	user.Password = changePass.Password
	if err := repository.GetUserUserByIDPass(&user); err != nil {
		if err.Error() == "record not found" {
			return http.StatusUnauthorized, err
		}
		return http.StatusInternalServerError, err
	}
	user.Password = changePass.NewPassword
	if err := repository.UpdateUser(user); err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusNoContent, nil
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
func ValidateJWT(Token string) (float64, float64, int, error) {
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
		return claims["ID"].(float64), claims["TypeID"].(float64), http.StatusOK, nil
	}
	return 0, 0, http.StatusInternalServerError, err
}
