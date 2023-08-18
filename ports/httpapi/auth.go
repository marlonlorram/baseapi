package httpapi

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/marlonlorram/baseapi/entities"
	"github.com/marlonlorram/baseapi/usecases/auth"
)

type Auth struct {
	service *auth.Service
}

type userInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// login autentica um usuário e retorna um token ou um erro.
func (a *Auth) login(ctx *gin.Context) (interface{}, error) {
	input := &userInput{}
	err := ctx.BindJSON(input)
	if err != nil {
		return nil, entities.NewBadRequest(fmt.Sprintf("Formato JSON inválido: %s", err.Error()))
	}

	return a.service.Login(ctx.Request.Context(), input.Username, input.Password)
}

// register cria um novo usuário e retorna informações sobre ele ou um erro.
func (a *Auth) register(ctx *gin.Context) (interface{}, error) {
	input := &userInput{}
	err := ctx.BindJSON(input)
	if err != nil {
		return nil, entities.NewBadRequest(fmt.Sprintf("Formato JSON inválido: %s", err.Error()))
	}

	return a.service.Register(ctx.Request.Context(), input.Username, input.Password)
}

// Mount define as rotas de autenticação.
func (a *Auth) Mount(router *gin.RouterGroup) {
	router.POST("/login", handler(a.login))
	router.POST("/register", handler(a.register))
}

// NewAuth cria uma nova instância do controlador Auth.
func NewAuth(service *auth.Service) *Auth {
	return &Auth{service: service}
}
