package httpapi

import (
	"github.com/gin-gonic/gin"

	"github.com/marlonlorram/baseapi"
)

// Version representa a estrutura do controlador da versão.
type Version struct{}

// getVersion obtém a versão atual da aplicação.
func (v *Version) getVersion(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "versão obtida",
		"data":    baseapi.Version,
	})
}

// Mount associa o controlador ao grupo de rotas.
func (v *Version) Mount(router *gin.RouterGroup) {
	router.GET("/version", v.getVersion)
}

// Cria um novo controlador para a versão.
func NewVersion() Version { return Version{} }
