package repository

import (
	"context"
	"errors"
	"time"

	"github.com/Stealthhy7512/gophercises/urlshort/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var ErrNotFound = errors.New("url not found")

type URLRepository interface {
	Insert(c context.Context, u *model.URL) (*model.URL, error)
	GetShortURLByLong(c context.Context, longURL string) (string, error)
	GetLongURLByShort(c context.Context, shortURL string) (string, error)
	UpdateShortURL(c context.Context, id bson.ObjectID, shortURL string) error
}

type mongoURLRepository struct {
	collection *mongo.Collection
}

func NewMongoURLRepository(c *mongo.Collection) (URLRepository, error) {
	r := &mongoURLRepository{
		collection: c,
	}

	if err := r.createIndexes(context.Background()); err != nil {
		return nil, err
	}

	return r, nil
}

func (r *mongoURLRepository) Insert(c context.Context, u *model.URL) (*model.URL, error) {
	u.CreatedAt = time.Now()

	res, err := r.collection.InsertOne(c, u)
	if err != nil {
		return nil, err
	}
	u.ID = res.InsertedID.(bson.ObjectID)

	return u, nil
}

func (r *mongoURLRepository) GetShortURLByLong(c context.Context, longURL string) (string, error) {
	res := model.URL{}

	err := r.collection.
		FindOne(c, bson.M{"long_url": longURL}).
		Decode(&res)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return "", ErrNotFound
		}
		return "", err
	}

	return res.ShortURL, nil
}

func (r *mongoURLRepository) GetLongURLByShort(c context.Context, shortURL string) (string, error) {
	var res model.URL
	err := r.collection.FindOne(c, bson.M{"short_url": shortURL}).Decode(&res)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return "", ErrNotFound
		}
		return "", err
	}
	return res.LongURL, nil
}

func (r *mongoURLRepository) UpdateShortURL(
	c context.Context, id bson.ObjectID, shortURL string,
) error {
	_, err := r.collection.UpdateByID(
		c,
		id,
		bson.M{
			"$set": bson.M{"short_url": shortURL},
		},
	)

	return err
}

func (r *mongoURLRepository) createIndexes(c context.Context) error {
	idxs := []mongo.IndexModel{
		{
			Keys:    bson.M{"long_url": 1},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys:    bson.M{"short_url": 1},
			Options: options.Index().SetUnique(true),
		},
	}
	_, err := r.collection.Indexes().CreateMany(c, idxs)

	return err
}
