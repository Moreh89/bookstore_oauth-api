package rest

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	fmt.Println("about to start test cases...")
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestLoginUserTimeoutFromApi(t *testing.T) {
	// TODO test not working due to err in external library github.com/mercadolibre/golang-restclient/rest
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.bookstore.com/users/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email":"email@gmail.com", "password":"password"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message": "invalid login credentials", "status": "forced_error", "error": "not_found"}`,
	})
	repository := usersRepository{}
	user, err := repository.LoginUser("email@gmail.com", "password")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid rest client response when trying to login user", err.Message)
}

// func TestLoginUserInvalidErrorInterface(t *testing.T)   {}
// func TestLoginUserInvalidLoginCredentials(t *testing.T) {}
// func TestLoginUserInvalidUserJsonResponse(t *testing.T) {}
func TestLoginUserNoError(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.bookstore.com/users/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email":"email@gmail.com", "password":"password"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"id": 1, "first_name": "damian", "last_name": "more", "email": "asd@asd.com"}`,
	})
	repository := usersRepository{}

	user, err := repository.LoginUser("email@gmail.com", "password")
	assert.NotNil(t, user)
	assert.Nil(t, err)
	assert.EqualValues(t, 1, user.Id)
	assert.EqualValues(t, "damian", user.FirstName)
	assert.EqualValues(t, "more", user.LastName)
	assert.EqualValues(t, "asd@asd.com", user.Email)

}
