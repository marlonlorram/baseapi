package usecases

import (
	"go.uber.org/fx"

	"github.com/marlonlorram/baseapi/usecases/auth"
	"github.com/marlonlorram/baseapi/usecases/users"
)

func Build() fx.Option {
	return fx.Options(
		auth.Build(),
		users.Build(),
	)
}
