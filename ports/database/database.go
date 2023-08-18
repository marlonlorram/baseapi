package database

import (
	"github.com/marlonlorram/baseapi/ports/database/mongodb"
	"github.com/marlonlorram/baseapi/ports/database/mongodb/users"
	"go.uber.org/fx"
)

func Build() fx.Option {
	return fx.Options(
		mongodb.Build(),
		users.Build(),
	)
}
