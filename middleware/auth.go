package middleware

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kevinjuliow/dataInventarisBarang/helper"
	"github.com/kevinjuliow/dataInventarisBarang/model/dtos"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path == "/api/register" || request.URL.Path == "/api/login" {
		middleware.Handler.ServeHTTP(writer, request)
		return
	}

	authorizationHeader := request.Header.Get("Authorization")
	if !strings.Contains(authorizationHeader, "Bearer") {
		middleware.unauthorized(writer, request)
		return
	}

	tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

	jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		middleware.unauthorized(writer, request)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		middleware.unauthorized(writer, request)
		return
	}

	userId := int(claims["user_id"].(float64))

	ctx := context.WithValue(request.Context(), "userId", userId)
	middleware.Handler.ServeHTTP(writer, request.WithContext(ctx))
}

func (middleware *AuthMiddleware) unauthorized(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusUnauthorized)

	webResponse := dtos.WebResponse{
		Code:   http.StatusUnauthorized,
		Status: "UNAUTHORIZED",
	}

	helper.WriteToResponseBody(writer, webResponse)
}
