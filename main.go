package main

import (
	"golang-restful-api/helper"
	"golang-restful-api/middleware"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func NewServer(newAuthMiddleware *middleware.AuthMiddleware) *http.Server{
	return &http.Server{
		Addr: "localhost:8080",
		Handler: newAuthMiddleware,
	}
}

func main() {
	server := InitializedServer()

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}