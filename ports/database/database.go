package database

import (
	"github.com/marlonlorram/baseapi/ports/database/mongodb"
	"go.uber.org/fx"
)

func Build() fx.Option {
	return fx.Options(
		mongodb.Build(),
	)
}
