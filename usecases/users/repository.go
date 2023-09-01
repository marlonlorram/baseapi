package users

import (
	"context"

	"github.com/marlonlorram/baseapi/entities"
)

type Repository interface {
	FindName(ctx context.Context, name string) (*entities.User, error)
	List(ctx context.Context) ([]*entities.User, error)
	Insert(ctx context.Context, user *entities.User) (*entities.User, error)
}
