package db

import (
	"context"

	"github.com/HsiaoCz/twitter-go/types"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostStorer interface {
	CreatePost(context.Context, *types.Posts) (*types.Posts, error)
}

type MongoPostStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func MongoPostStoreInit(client *mongo.Client, coll *mongo.Collection) *MongoPostStore {
	return &MongoPostStore{
		client: client,
		coll:   coll,
	}
}

func (m *MongoPostStore) CreatePost(ctx context.Context, post *types.Posts) (*types.Posts, error) {
	return nil, nil
}
