package app

import (
	accesstoken "Gone/src/domain/access_token"
	"Gone/src/http"
	"Gone/src/repository/db"
)

func StartApplication(){
	dbRepository := db.NewRepository()
	atService := accesstoken.NewService(dbRepository)
	atHandler := http.NewHandler(atService)
	
	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.Run(":8080")
}

