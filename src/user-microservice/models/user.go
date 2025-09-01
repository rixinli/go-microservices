/*
 * @File: models.user.go
 * @Description: Defines User model
 * @Author: Nguyen Truong Duong (seedotech@gmail.com)
 */
package models

import (
	"errors"

	"github.com/rixinli/go-microservices/src/user-microservice/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User information
type User struct {
	ID       primitive.ObjectID `bson:"_id" json:"id" example:"5bbdadf782ebac06a695a8e7"`
	Name     string        `bson:"name" json:"name" example:"raycad"`
	Password string        `bson:"password" json:"password" example:"raycad"`
}
// Validate user
func (a AddUser) Validate() error {
	switch {
	case len(a.Name) == 0:
		return errors.New(common.ErrNameEmpty)
	case len(a.Password) == 0:
		return errors.New(common.ErrPasswordEmpty)
	default:
		return nil
	}
}

// AddUser information
type AddUser struct {
	Name     string `json:"name" example:"User Name"`
	Password string `json:"password" example:"User Password"`
}

type UpdateUser struct {
	ID       primitive.ObjectID `bson:"_id" json:"id" example:"5bbdadf782ebac06a695a8e7"`
	Name     string `json:"name,omitempty" example:"User Name"`
	Password string `json:"password,omitempty" example:"User Password"`
}

// Validate checks if fields are empty and returns error if so
func (u UpdateUser) Validate() error {
	switch {
	case len(u.Name) == 0:
		return errors.New(common.ErrNameEmpty)
	case len(u.Password) == 0:
		return errors.New(common.ErrPasswordEmpty)
	default:
		return nil
	}
}



