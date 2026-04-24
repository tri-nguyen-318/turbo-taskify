package mongo

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"

	"turbo-taskify/backend/internal/domain/user"
)

type UserRepository struct {
	col *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	col := db.Collection("users")

	col.Indexes().CreateMany(context.Background(), []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "email", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys:    bson.D{{Key: "username", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	})

	return &UserRepository{col: col}
}

func (r *UserRepository) Create(ctx context.Context, u *user.User) error {
	_, err := r.col.InsertOne(ctx, u)
	if mongo.IsDuplicateKeyError(err) {
		return user.ErrEmailExists
	}
	return err
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*user.User, error) {
	var u user.User
	err := r.col.FindOne(ctx, bson.M{"email": email}).Decode(&u)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, user.ErrNotFound
	}
	return &u, err
}

func (r *UserRepository) FindByID(ctx context.Context, id string) (*user.User, error) {
	oid, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, user.ErrNotFound
	}
	var u user.User
	err = r.col.FindOne(ctx, bson.M{"_id": oid}).Decode(&u)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, user.ErrNotFound
	}
	return &u, err
}

func (r *UserRepository) FindByGoogleID(ctx context.Context, googleID string) (*user.User, error) {
	var u user.User
	err := r.col.FindOne(ctx, bson.M{"google_id": googleID}).Decode(&u)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, user.ErrNotFound
	}
	return &u, err
}

func (r *UserRepository) ExistsEmail(ctx context.Context, email string) (bool, error) {
	n, err := r.col.CountDocuments(ctx, bson.M{"email": email})
	return n > 0, err
}

func (r *UserRepository) ExistsUsername(ctx context.Context, username string) (bool, error) {
	n, err := r.col.CountDocuments(ctx, bson.M{"username": username})
	return n > 0, err
}
