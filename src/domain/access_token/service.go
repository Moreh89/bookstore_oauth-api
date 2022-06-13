package accesstoken

import (
	"Gone/src/utils/errors"
)

type Service interface {
	GetById(string) (*AccessToken, *errors.RestError)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetById(string) (*AccessToken, *errors.RestError) {
	return nil, nil
}
