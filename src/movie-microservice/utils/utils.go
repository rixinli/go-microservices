/*
 * @File: utils.utils.go
 * @Description: Reusable stuffs for services
 * @Author: Nguyen Truong Duong (seedotech@gmail.com)
 */
package utils

import (
	"errors"
	"time"

	jwt_lib "github.com/dgrijalva/jwt-go"
	"github.com/rixinli/go-microservices/src/movie-microservice/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SdtClaims defines the custom claims
type SdtClaims struct {
	Name string `json:"name"`
	Role string `json:"role"`
	jwt_lib.StandardClaims
}

type Utils struct {
}

// GenerateJWT generates token from the given information
func (u *Utils) GenerateJWT(name string, role string) (string, error) {
	claims := SdtClaims{
		name,
		role,
		jwt_lib.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			Issuer:    common.Config.Issuer,
		},
	}

	token := jwt_lib.NewWithClaims(jwt_lib.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(common.Config.JwtSecretPassword))

	return tokenString, err
}

// ValidateObjectID checks the given ID if it's an object id or not
func (u *Utils) ValidateObjectID(id string) error {
	_, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return errors.New(common.ErrNotObjectIDHex)
    }
    return nil
}
