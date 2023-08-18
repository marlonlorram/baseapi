package usecases

import (
	"go.uber.org/fx"

	"github.com/marlonlorram/baseapi/usecases/auth"
)

func Build() fx.Option {
	return fx.Options(
		auth.Build(),
	)
}
