package ports

import (
	"go.uber.org/fx"

	"github.com/marlonlorram/baseapi/ports/httpapi"
)

func Build() fx.Option {
	return fx.Provide(
		httpapi.NewVersion,
	)
}
