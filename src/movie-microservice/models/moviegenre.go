/*
 * @File: models.movie_genre.go
 * @Description: Defines Movie Genre information will be returned to the clients
 * @Author: Nguyen Truong Duong (seedotech@gmail.com)
 */
package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// MovieGenre information
type MovieGenre struct {
	ID          primitive.ObjectID `bson:"id" json:"id"`
	Name        string        `bson:"name" json:"name"`
	Description string        `bson:"description" json:"description"`
}
