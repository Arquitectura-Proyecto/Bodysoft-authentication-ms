package repository

import (
	"strconv"
	"time"
	"fmt"
	"context"
	guuid "github.com/google/uuid"
	"github.com/jpbmdev/Bodysoft-authentication-ms/data"
	"github.com/jpbmdev/Bodysoft-authentication-ms/models"
	"github.com/jpbmdev/Bodysoft-authentication-ms/utils"
)

var ctx = context.Background()
// FindUserByEmail ..
func FindUserByEmail(user models.User) error {
	db := data.DatabaseConection()
	defer db.Close()
	if err := db.Where("Email = ?", user.Email).First(&user).Error; err != nil {
		return err
	}
	return nil
}

// GetUserByEmail ..
func GetUserByEmail(user *models.User) error {
	db := data.DatabaseConection()
	defer db.Close()
	if err := db.Where("Email = ?", user.Email).First(&user).Error; err != nil {
		return err
	}
	return nil
}

// GetUserByEmailPass ..
func GetUserByEmailPass(user *models.User) error {

	db := data.DatabaseConection()
	defer db.Close()
	if err := db.Where("Email = ? AND Password = ?", user.Email, user.Password).First(&user).Error; err != nil {
		return err
	}
	return nil
}

// GetUserByID ..
func GetUserByID(user *models.User) error {
	db := data.DatabaseConection()
	defer db.Close()
	if err := db.Where("ID = ?", user.ID).First(&user).Error; err != nil {
		return err
	}
	return nil
}

// GetUserByIDPass ..
func GetUserByIDPass(user *models.User) error {
	db := data.DatabaseConection()
	defer db.Close()
	if err := db.Where("ID = ? AND Password = ?", user.ID, user.Password).First(&user).Error; err != nil {
		return err
	}
	return nil
}

// CreateUser .
func CreateUser(user models.User) error {
	db := data.DatabaseConection()
	defer db.Close()
	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

// UpdateUser ..
func UpdateUser(user models.User) error {
	db := data.DatabaseConection()
	defer db.Close()
	if err := db.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUserPassword(user models.ChangePass) error {
	db := data.DatabaseConection()
	rdb := data.RedisDbConection()
	defer db.Close()
	defer rdb.Close()
	fmt.Println(user.Token)
	email,err := rdb.Get(ctx,user.Token).Result()
	if err!= nil{
		fmt.Println(err)
	}
	var u models.User
	u.Email = email
	if err=GetUserByEmail(&u); err!=nil{
		return err
	}

	if err := db.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

// GenerateRecoveryEmail .
func GenerateRecoveryEmail(email string) error {
	rdb := data.RedisDbConection()
	defer rdb.Close()
	uid := guuid.New().String()
	err := rdb.Set(ctx,uid, email, time.Hour*60).Err()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("token",uid)
	subject := "Cambio de Contraseña Bodysoft"
	url:= "https://bodysoft.me/newpassword/"+ uid
	htmlContent:= utils.EmailTemplate(url)
	if err := data.SendEmail(email, htmlContent, subject); err != nil {
		return err
	}
	return nil
}

func IsChangeValid(uid string) (string, error){
	rdb := data.RedisDbConection()
	defer rdb.Close()
	fmt.Println(uid)
	email, err := rdb.Get(ctx, uid).Result()
	if err != nil {
		return "",err
	}
	rdb.Del(ctx,uid)
	return email,nil
}

// GenerateValidationEmail .
func GenerateValidationEmail(email string, vcode uint) error {
	subject := "Código de Verification BodySoft"
	htmlContent := "<h1> Su codigo es: " + strconv.Itoa(int(vcode)) + "</h1>"
	if err := data.SendEmail(email, htmlContent, subject); err != nil {
		return err
	}
	return nil
}
