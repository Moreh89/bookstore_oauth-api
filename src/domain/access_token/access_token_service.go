package accesstoken

import (
	"Gone/src/utils/errors"
	"fmt"
	"strings"
)

type Service interface {
	GetById(string) (*AccessToken, *errors.RestError)
	Create(AccessToken) *errors.RestError
	UpdateExpirationTime(AccessToken) *errors.RestError
}

type Repository interface {
	GetById(string) (*AccessToken, *errors.RestError)
	Create(AccessToken) *errors.RestError
	UpdateExpirationTime(AccessToken) *errors.RestError
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
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		fmt.Println(accessTokenId)
		return nil, errors.NewBadRequestError("invalid access token id")
	}
	accessToken, err := s.repository.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) Create(req AccessTokenRequest) (*AccessToken ,*errors.RestError) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	user, err := s.repository.LoginUser(req.Username, req.Password)
	if err != nil {
		return nil, err
	}
	at := AccessToken.GetNewAccessToken(user.Id)
	at.Generate
}

// func (s *service) Create(at AccessToken) *errors.RestError {
// 	err := at.Validate(); if err != nil {
// 		return err
// 	}

// 	return s.repository.Create(at)
// }

func (s *service) UpdateExpirationTime(at AccessToken) *errors.RestError {
	err := at.Validate(); if err != nil {
		return err
	}
	return s.repository.UpdateExpirationTime(at)
}
