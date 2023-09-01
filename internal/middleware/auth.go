package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/marlonlorram/baseapi/usecases/auth"
)

type AuthMiddleware struct {
	service *auth.Service
}

// Authorization é um middleware para autenticação JWT.
// Ele interrompe a requisição se o token JWT não for válido.
func (am *AuthMiddleware) Authorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		abortRequest := func() {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "Acesso não autorizado!",
			})
		}

		header := ctx.Request.Header.Get("Authorization")
		if header == "" {
			abortRequest()
			return
		}

		splitToken := strings.Split(header, "Bearer ")
		if len(splitToken) != 2 {
			abortRequest()
			return
		}

		user, err := am.service.ValidateToken(splitToken[1])
		if err != nil {
			abortRequest()
			return
		}

		ctx.Set("id", user.ID)
		ctx.Set("name", user.Name)
		ctx.Next()
	}
}

// newAuthMiddleware cria uma nova instância do middleware de autenticação.
func newAuthMiddleware(service *auth.Service) *AuthMiddleware {
	return &AuthMiddleware{service: service}
}
