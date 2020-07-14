package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"meetmeup/graph/model"
)

type UsersRepo struct {
	//DB *mongo.Client
}

func (u *UsersRepo) GetUserByField(field, value string) (*model.User, error) {
	var user model.User
	err := UserCollection.FindOne(context.TODO(), bson.M{field: value}).Decode(&user)
	return &user, err
}

func (u *UsersRepo) GetUserById(id string) (*model.User, error) {
	return u.GetUserByField("_id", id)
}

func (u *UsersRepo) GetUserByEmail(email string) (*model.User, error) {
	return u.GetUserByField("email", email)
}

func (u *UsersRepo) GetUserByUsername(username string) (*model.User, error) {
	return u.GetUserByField("username", username)
}

func (u *UsersRepo) CreateUser(user *model.User) (*model.User, error) {
	_, err := UserCollection.InsertOne(context.TODO(), user)
	return user, err
}
