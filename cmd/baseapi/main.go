package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/marlonlorram/baseapi/config"
	"github.com/marlonlorram/baseapi/ports"
	"github.com/marlonlorram/baseapi/ports/database"
	"github.com/marlonlorram/baseapi/server"
)

// newApp inicializa uma nova instância da aplicação fx com as configurações e
// dependências necessárias.
func newApp(ctx context.Context) *fx.App {
	return fx.New(
		fx.Provide(func() context.Context { return ctx }),

		fx.Provide(func() (*zap.Logger, error) {
			return zap.NewProduction()
		}),

		config.Build(),
		database.Build(),
		ports.Build(),
		server.Build(),
	)
}

// start inicializa e gerencia o ciclo de vida da aplicação.
func start(ctx context.Context) {
	app := newApp(ctx)

	if err := app.Start(ctx); err != nil {
		zap.L().Error("erro ao iniciar a aplicação", zap.Error(err))
		os.Exit(1)
	}

	// Aguarda até que o contexto da aplicação seja finalizado.
	<-ctx.Done()

	// Estabelece um contexto com um timeout de 15 segundos para permitir que a aplicação
	// encerre suas operações de forma controlada. Se a aplicação não parar dentro deste
	// período, o contexto será cancelado, forçando a aplicação a terminar.
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := app.Stop(ctx); err != nil {
		zap.L().Error("erro ao tentar parar a aplicação", zap.Error(err))
		os.Exit(1)
	}
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	start(ctx)
}
