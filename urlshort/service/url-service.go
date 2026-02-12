package service

import (
	"context"
	"time"

	"github.com/Stealthhy7512/gophercises/urlshort/model"
	"github.com/Stealthhy7512/gophercises/urlshort/repository"
	"github.com/jxskiss/base62"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type URLService interface {
	ShortenURL(c context.Context, longURL string) (string, error)
	GetLongURL(c context.Context, shortURL string) (string, error)
}

type urlService struct {
	r repository.URLRepository
}

func NewURLService(r repository.URLRepository) URLService {
	return &urlService{
		r: r,
	}
}

func (s *urlService) ShortenURL(c context.Context, longURL string) (string, error) {
	u := &model.URL{
		LongURL:   longURL,
		CreatedAt: time.Now(),
	}

	inserted, err := s.r.Insert(c, u)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			existing, err := s.r.GetShortURLByLong(c, longURL)
			if err != nil {
				return "", err
			}
			return existing, nil
		}
		return "", err
	}

	shortCode := base62.EncodeToString(inserted.ID[:4])

	err = s.r.UpdateShortURL(c, inserted.ID, shortCode)
	if err != nil {
		return "", err
	}

	return shortCode, nil
}

func (s *urlService) GetLongURL(c context.Context, shortURL string) (string, error) {
	return s.r.GetLongURLByShort(c, shortURL)
}
