package users

import (
	"context"

	"go.uber.org/fx"

	"github.com/marlonlorram/baseapi/entities"
)

type Service struct {
	repo Repository
}

// GetUsers obtém todos os usuários registrados.
func (s *Service) GetUsers(ctx context.Context) ([]*entities.User, error) {
	return s.repo.List(ctx)
}

// newUserService cria uma nova instância do serviço de usuários.
func newUserService(repository Repository) *Service {
	return &Service{repository}
}

func Build() fx.Option {
	return fx.Provide(newUserService)
}
