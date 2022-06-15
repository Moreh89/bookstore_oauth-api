package app

import (
	accesstoken "Gone/src/domain/access_token"
	"Gone/src/http"
	"Gone/src/repository/db"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	dbRepository := db.NewRepository()
	atService := accesstoken.NewService(dbRepository)
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)

	router.Run(":8080")
}
