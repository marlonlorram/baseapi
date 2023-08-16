package httpapi

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/marlonlorram/baseapi/entities"
)

// response representa a estrutura padrão de resposta para as solicitações HTTP.
type response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// handler é um função auxiliar que padroniza a forma como as respostas HTTP são retornadas.
func handler(fun func(c *gin.Context) (interface{}, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		val, err := fun(c)
		if err != nil {
			errCode := newError(err)

			c.JSON(errCode, response{
				Code:    errCode,
				Message: err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, response{
			Code:    http.StatusOK,
			Message: "sucesso",
			Data:    val,
		})
	}
}

// newError mapeia um erro para um código de status HTTP apropriado.
func newError(err error) int {
	switch {
	case entities.IsNotFound(err):
		return http.StatusNotFound
	case entities.IsUnauthorized(err):
		return http.StatusUnauthorized
	case entities.IsForbidden(err):
		return http.StatusForbidden
	case entities.IsBadRequest(err):
		return http.StatusBadRequest
	case entities.IsConflict(err):
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
