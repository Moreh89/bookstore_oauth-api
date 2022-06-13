package http

import (
	accesstoken "Gone/src/domain/access_token"
	"Gone/src/utils/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccessTokenHandler interface {
	GetById(c *gin.Context)
}

type accessTokenHandler struct {
	service accesstoken.Service
}

func NewHandler(service accesstoken.Service) {
	return &accessTokenHandler{
		service: service,
	}
}

func (h *accessTokenHandler) GetById(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, "implement me!")
}
