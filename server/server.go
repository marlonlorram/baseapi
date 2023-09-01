package server

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/marlonlorram/baseapi/config"
	"github.com/marlonlorram/baseapi/internal/middleware"
	"github.com/marlonlorram/baseapi/ports/httpapi"
)

// Bindings encapsula as dependências que serão injetadas
type Bindings struct {
	fx.In
	Version httpapi.Version
	Auth    *httpapi.Auth
}

// newRouter retorna uma instância Gin com as rotas predefinidas.
func newRouter(b Bindings) *gin.Engine {
	router := gin.Default()

	router.Use(middleware.NewCorsMiddleware())

	apiV1 := router.Group("/api/v1")

	b.Version.Mount(apiV1)
	b.Auth.Mount(apiV1)
	return router
}

// newServer inicializa o servidor HTTP usando o framework Gin e configura
// ganchos para gerenciar seu ciclo de vida através do fx.Lifecycle.
func newServer(lc fx.Lifecycle, l *zap.Logger, cfg *config.Config, router *gin.Engine) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			l.Info("iniciando servidor http")
			go func() {
				if err := router.Run(cfg.Server.Address); err != nil {
					l.Fatal("falha ao iniciar o servidor http", zap.Error(err))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			l.Info("parando servidor http")
			return nil
		},
	})
}

func Build() fx.Option {
	return fx.Options(
		fx.Provide(newRouter),
		fx.Invoke(newServer),
	)
}
