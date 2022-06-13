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

func (s *service) GetById(accessTokenId string) (*AccessToken, *errors.RestError) {
	accessToken, err := s.repository.GetById((accessTokenId))
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}
