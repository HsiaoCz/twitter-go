package db

import (
	"context"

	"github.com/HsiaoCz/twitter-go/types"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserStorer interface {
	CreateUser(context.Context, *types.CreateUserParams) (*types.User, error)
}

type MongoUserStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoUserStore(client *mongo.Client, coll *mongo.Collection) *MongoUserStore {
	return &MongoUserStore{
		client: client,
		coll:   coll,
	}
}

func (m *MongoUserStore) CreateUser(ctx context.Context, params *types.CreateUserParams) (*types.User, error) {
	return nil,nil
}
