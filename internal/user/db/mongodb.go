package db

import (
	"context"
	"fmt"
	"test_RESTserver_01/internal/user"
	"test_RESTserver_01/pkg/logging"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type db struct {
	collection	*mongo.Collection
	logger		*logging.Logger
}

// Create implements user.Storage
func (d *db) Create(ctx context.Context, user user.User) (string, error) {
	d.logger.Debug("create user")
	result, err := d.collection.InsertOne(ctx, user)
	if err != nil {
		return "", fmt.Errorf("error create user: %v", err)
	}

	d.logger.Debug("conver insertedID to objectID")
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return oid.Hex(), nil
	}
	d.logger.Trace(user)
	return "", fmt.Errorf("failed to convert objectid to hex")
}

// Delete implements user.Storage
func (*db) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}

// FindOne implements user.Storage
func (*db) FindOne(ctx context.Context, id string) (user.User, error) {
	panic("unimplemented")
}

// Update implements user.Storage
func (*db) Update(ctx context.Context, user user.User) error {
	panic("unimplemented")
}

func NewStorage(database *mongo.Database, collection string, logger *logging.Logger) user.Storage {
	return &db{
		collection: database.Collection(collection),
		logger: logger,
	}
}
