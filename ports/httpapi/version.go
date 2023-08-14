package httpapi

import (
	"github.com/gin-gonic/gin"

	"github.com/marlonlorram/baseapi"
)

// Version representa a estrutura do controlador da versão.
type Version struct{}

// getVersion obtém a versão atual da aplicação.
func (v *Version) getVersion(ctx *gin.Context) (interface{}, error) {
	return baseapi.Version, nil
}

// Mount associa o controlador ao grupo de rotas.
func (v *Version) Mount(router *gin.RouterGroup) {
	router.GET("/version", handler(v.getVersion))
}

// Cria um novo controlador para a versão.
func NewVersion() Version { return Version{} }
