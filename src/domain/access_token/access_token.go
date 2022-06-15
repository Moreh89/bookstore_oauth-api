package accesstoken

import (
	"Gone/src/utils/errors"
	"fmt"
	"strings"
	"time"
)

const (
	expirationTime = 24
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func (at *AccessToken) Validate() *errors.RestError {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		fmt.Println("access token")
		fmt.Println(at.AccessToken)
		return errors.NewBadRequestError("invalid access token id")
	}
	if at.UserId < 0 {
		return errors.NewBadRequestError("invalid user id")
	}
	if at.ClientId < 0 {
		return errors.NewBadRequestError("invalid client id")
	}
	if at.Expires < 0 {
		return errors.NewBadRequestError("invalid expiration time")
	}
	return nil
}

// type Repository interface {
// 	GetById(string) (*AccessToken, *errors.RestError)
// }

func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}
