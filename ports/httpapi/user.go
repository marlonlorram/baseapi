package httpapi

import (
	"github.com/gin-gonic/gin"

	"github.com/marlonlorram/baseapi/internal/middleware"
	"github.com/marlonlorram/baseapi/usecases/users"
)

type User struct {
	service        *users.Service
	authMiddleware *middleware.AuthMiddleware
}

// getUsers obtém todos os usuários registrados.
func (u *User) getUsers(ctx *gin.Context) (interface{}, error) {
	return u.service.GetUsers(ctx.Request.Context())
}

// Mount define as rotas de usuário.
func (u *User) Mount(router *gin.RouterGroup) {
	router.Use(u.authMiddleware.Authorization())

	router.POST("/users", handler(u.getUsers))
}

// NewUser cria uma nova instância do controlador User.
func NewUser(service *users.Service, authMiddleware *middleware.AuthMiddleware) *User {
	return &User{service, authMiddleware}
}