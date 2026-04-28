package repository

import (
	"context"
	"errors"
	"time"

	"github.com/agdaniel10/Go-BasicAPI/internal/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{
		collection: db.Collection("users"),
	}
}

func (r *UserRepository) Create(ctx context.Context, user *model.User) error {
	user.CreatedAt = time.Now()
	_, err := r.collection.InsertOne(ctx, user)
	return err
}

func (r *UserRepository) FindByID(ctx context.Context, id string) (*model.User, error) {
	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var user model.User
	err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var user model.User
	err := r.collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)

	if err == mongo.ErrNoDocuments {
		return nil, nil
	}

	if err != nil {
		return nil, nil
	}
	return &user, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, id string, user *model.User) (*model.User, error) {

	if id == "" {
		return nil, errors.New("id is required")
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	updateFields := bson.M{}

	if user.Name != "" {
		updateFields["name"] = user.Name
	}

	if user.Email != "" {
		updateFields["email"] = user.Email
	}

	if len(updateFields) == 0 {
		return nil, errors.New("No fields found")
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": updateFields}

	ops := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var updatedUser model.User
	err = r.collection.FindOneAndUpdate(ctx, filter, update, ops).Decode(&updatedUser)

	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &updatedUser, nil
}

func (r *UserRepository) FindAll(ctx context.Context) ([]model.User, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var users []model.User

	if err := cursor.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}
