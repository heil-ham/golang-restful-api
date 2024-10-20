package middleware

import (
	"golang-restful-api/helper"
	"golang-restful-api/model/web"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware{
	return &AuthMiddleware{
		Handler: handler,
	}
}

func (authMiddleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Header.Get("X-API-Key") == "RAHASIA" {
		authMiddleware.Handler.ServeHTTP(writer, request)
	} else {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		helper.WriteToResponseBody(writer, webResponse)	
	}
}

