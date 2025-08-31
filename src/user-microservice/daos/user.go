/*
 * @File: daos.user.go
 * @Description: Implements User CRUD functions for MongoDB
 * @Author: Nguyen Truong Duong (seedotech@gmail.com)
 */
package daos

import (
	"context"
	"errors"
	"time"

	"github.com/rixinli/go-microservices/src/user-microservice/common"
	"github.com/rixinli/go-microservices/src/user-microservice/databases"
	"github.com/rixinli/go-microservices/src/user-microservice/models"
	"github.com/rixinli/go-microservices/src/user-microservice/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User manages User CRUD
type User struct {
	utils *utils.Utils
}

// GetAll gets the list of Users
func (u *User) GetAll() ([]models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	// get all ColUser collections
	collection := databases.Database.Client.Database(databases.Database.Databasename).Collection(common.ColUsers)

	var users []models.User

	// cursor and error 
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil{
		return nil,err
	}
	defer cursor.Close(ctx)

	// writing all users from cursor into users array
	if err = cursor.All(ctx,&users); err !=nil{
		return nil,err
	}

	return users, err
}

// GetByID finds a User by its id
func (u *User) GetByID(id string) (models.User, error) {
	var err error
	err = u.utils.ValidateObjectID(id)
	if err != nil {
		return models.User{}, err
	}

	// transform into primitive.ObjectID
    oid, _ := primitive.ObjectIDFromHex(id)

	// query the collection and user
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := databases.Database.Client.Database(databases.Database.Databasename).Collection(common.ColUsers)

	var user models.User
	err = collection.FindOne(ctx,bson.M{"_id":oid}).Decode(&user)

	if err!=nil {
		return models.User{},err
	}

	return user, nil
}

// DeleteByID finds a User by its id
func (u *User) DeleteByID(id string) error {
	// 验证 ObjectID
    if err := u.utils.ValidateObjectID(id); err != nil {
        return err
    }
    oid, _ := primitive.ObjectIDFromHex(id)

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    collection := databases.Database.Client.
        Database(databases.Database.Databasename).
        Collection(common.ColUsers)

	// delete
    res, err := collection.DeleteOne(ctx, bson.M{"_id": oid})
    if err != nil {
        return err
    }
    if res.DeletedCount == 0 {
        return errors.New("no document found with given ID")
    }
    return nil
}

// Login User
func (u *User) Login(name string, password string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Get a collection to execute the query against.
	collection := databases.Database.Client.
        Database(databases.Database.Databasename).
        Collection(common.ColUsers)

	var user models.User
	err := collection.FindOne(ctx, bson.M{"$and": []bson.M{
		{"name": name},
		{"password": password},
	}}).Decode(&user)
	return user, err
}

// Insert adds a new User into database'
func (u *User) Insert(user models.User) error {
	
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Get a collection to execute the query against.
	collection := databases.Database.Client.
        Database(databases.Database.Databasename).
        Collection(common.ColUsers)

	_, err := collection.InsertOne(ctx,&user)
	return err
}

// Delete remove an existing User
func (u *User) Delete(user models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Get a collection to execute the query against.
	collection := databases.Database.Client.
        Database(databases.Database.Databasename).
        Collection(common.ColUsers)

	_, err := collection.DeleteOne(ctx, bson.M{"_id": user.ID})
	return err
}

// Update modifies an existing User
func (u *User) Update(user models.User) error {
	

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Get a collection to execute the query against.
	collection := databases.Database.Client.
        Database(databases.Database.Databasename).
        Collection(common.ColUsers)

	_, err := collection.UpdateByID(ctx, user.ID, bson.M{"$set": user})
	return err
}
