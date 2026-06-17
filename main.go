package main

import (
	_ "emtesttask/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error *ErrorInfo  `json:"error,omitempty"`
}

type ErrorInfo struct {
	Message string `json:"message"`
}

// @title     Subscriptions Service
// @version         1.0
// @description     Effective Mobile test task
func main() {
	router, cleanup, _ := InitializeApp()
	defer cleanup()

	router.GET("/swagger-ui/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":8080")
}
