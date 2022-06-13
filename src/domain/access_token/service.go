package accesstoken

import (
	// "Gone/src/utils/errors"
	"github.com/Moreh89/bookstore_oauth-api/utils/errors"
)

type Service interface {
	GetById(string) (*AccessToken, *errors.RestError)
}

type service struct {
}

func NewService() Service {
	return &service{}
}

func (s *service) GetById(string) (*AccessToken, *errors.RestError) {
	return nil, nil
}