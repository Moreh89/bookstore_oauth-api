package accesstoken

import "github.com/Moreh89/bookstore_oauth-api/src/utils/errors"

type Service interface {
	GetById(string) (*AccessToken, *errors.RestErr)
}

type service struct {
}

func NewService() Service {
	return &service{}
}

func (s *service) GetById(string) (*AccessToken, *errors.RestErr) {
	return nil, nil
}