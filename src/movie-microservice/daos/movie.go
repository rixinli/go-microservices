/*
 * @File: daos.movie.go
 * @Description: Implements Movie CRUD functions for MongoDB
 * @Author: Nguyen Truong Duong (seedotech@gmail.com)
 */
package daos

import (
	"context"
	"time"

	"github.com/rixinli/go-microservices/src/movie-microservice/databases"
	"github.com/rixinli/go-microservices/src/movie-microservice/models"
	"github.com/rixinli/go-microservices/src/movie-microservice/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Movie manages Movie CRUD
type Movie struct {
	utils *utils.Utils
}

const (
	ColMovies = "movies"
)

// GetAll gets the list of Movie
func (m *Movie) GetAll() ([]models.Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()


	// Get a collection of movie. ColMovies Collection
	collection := databases.Database.Client.Database(databases.Database.Databasename).Collection(ColMovies)

	var movies []models.Movie

	// cursor and error
	cursor, err := collection.Find(ctx,bson.M{})
	if err!=nil {
		return nil,err
	}
	
	// writing all movie into movies array
	if err = cursor.All(ctx,&movies); err!=nil {	
		return nil,err
	}
	
	return movies, err
}

// GetByID finds a Movie by its id
func (m *Movie) GetByID(id string) (models.Movie, error) {

	var err error
	err = m.utils.ValidateObjectID(id)
	if err != nil {
		return models.Movie{}, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// transform into primitive.ObjectID
    oid, _ := primitive.ObjectIDFromHex(id)
	
	// Get a collection to execute the query against.
	collection := databases.Database.Client.Database(databases.Database.Databasename).Collection(ColMovies)

	var movie models.Movie
	err = collection.FindOne(ctx,bson.M{"_id":oid}).Decode(&movie)
	return movie, err
}

// Insert adds a new Movie into database'
func (m *Movie) Insert(movie models.Movie) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Get a collection to execute the query against.
	collection := databases.Database.Client.
        Database(databases.Database.Databasename).
        Collection(ColMovies)

	_, err := collection.InsertOne(ctx,&movie)
	return err
}

// Delete remove an existing Movie
func (m *Movie) Delete(movie models.Movie) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Get a collection to execute the query against.
	collection := databases.Database.Client.
        Database(databases.Database.Databasename).
        Collection(ColMovies)

	_, err := collection.DeleteOne(ctx, bson.M{"_id": movie.ID})
	return err
}

// Update modifies an existing Movie
func (m *Movie) Update(movie models.Movie) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Get a collection to execute the query against.
	collection := databases.Database.Client.
        Database(databases.Database.Databasename).
        Collection(ColMovies)

	_, err := collection.UpdateByID(ctx, movie.ID, bson.M{"$set": movie})
	return err
}
