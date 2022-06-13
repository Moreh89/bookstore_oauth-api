package db

import (
	accesstoken "Gone/src/domain/access_token"
	"Gone/src/utils/errors"
)

type DbRepository interface {
	GetById(string) (*accesstoken.AccessToken, *errors.RestError)
}

type dbRepository struct {
}

func NewRepository() DbRepository {
	return &dbRepository{}
}

func (r *dbRepository) 	GetById(id string) (*accesstoken.AccessToken, *errors.RestError){
	return nil, nil
}