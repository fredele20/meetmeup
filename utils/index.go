package utils

//
//import (
//	"github.com/dgrijalva/jwt-go"
//	"golang.org/x/crypto/bcrypt"
//	"meetmeup/graph/model"
//	"time"
//)
//
//func (u model.User) HashPassword(password string) error {
//	bytePassword := []byte(password)
//	passwordHash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
//	if err != nil {
//		return err
//	}
//
//	u.Password = string(passwordHash)
//	return nil
//}
//func (u *User) GenToken() (*AuthToken, error) {
//	expiredAt := time.Now().Add(time.Hour * 24 * 7) // one week
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
//		ExpiresAt: expiredAt.Unix(),
//		Id:        u.ID,
//		IssuedAt:  time.Now().Unix(),
//		Issuer:    "meetmeup",
//	})
//}
//
