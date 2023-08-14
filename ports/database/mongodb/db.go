package mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/fx"

	"github.com/marlonlorram/baseapi/config"
)

// newClient retorna uma conexão ativa com o MongoDB. Utiliza as configurações
// especificadas em cfg. Se a porta não estiver especificada, a padrão 27018 é usada.
func newClient(lc fx.Lifecycle, cfg *config.Config) (*mongo.Database, error) {
	port := cfg.Database.Port
	if port <= 0 {
		port = 27018
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	credential := options.Credential{
		Username:      cfg.Database.User,
		Password:      cfg.Database.Pass,
		AuthSource:    cfg.Database.Name,
		AuthMechanism: cfg.Database.Mech,
	}

	clientOpts := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", cfg.Database.Host, port)).SetAuth(credential)

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, fmt.Errorf("falha ao conectar com o banco de dados: %w", err)
	}

	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		return nil, fmt.Errorf("falha ao verificar conexão com o banco de dados: %v", err)
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			fmt.Print("desconectando...")
			return client.Disconnect(ctx)
		},
	})

	db := client.Database(cfg.Database.Name)

	return db, nil
}

func Build() fx.Option {
	return fx.Provide(newClient)
}
