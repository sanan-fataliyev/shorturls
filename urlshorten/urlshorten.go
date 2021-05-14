package urlshorten

import (
	"fmt"
	"github.com/lithammer/shortuuid"
)

type Storage interface {
	GetOriginUrl(shortURL string) (originURL string, found bool)
	Save(originURL, shortURL string) error
}

type Service struct {
	baseURL string
	storage Storage
}

func NewService(baseURL string, storage Storage) *Service {
	return &Service{baseURL: baseURL, storage: storage}
}

func (s *Service) CreateShortURL(originURL string) (string, error) {
	short := shortuuid.New()
	shortURL := fmt.Sprintf("%s/%s", s.baseURL, short)
	err := s.storage.Save(originURL, shortURL)

	if err != nil {
		return "", err
	}
	return shortURL, nil

}

func (s *Service) GetOriginURL(shortURL string) (originURL string, found bool) {

	originURL, found = s.storage.GetOriginUrl(shortURL)
	return
}
