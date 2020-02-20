package typical

import (
	"github.com/typical-go/typical-rest-server/pkg/typdocker"
	"github.com/typical-go/typical-rest-server/pkg/typpostgres"
	"github.com/typical-go/typical-rest-server/pkg/typrails"
	"github.com/typical-go/typical-rest-server/pkg/typreadme"
	"github.com/typical-go/typical-rest-server/pkg/typredis"
	"github.com/typical-go/typical-rest-server/restserver"
)

// Modules that required for the project
var (
	rest     = restserver.New().WithDebug(true)
	readme   = typreadme.New()
	rails    = typrails.New()
	redis    = typredis.New()
	postgres = typpostgres.New().WithDBName("sample")

	docker = typdocker.New().WithComposers(
		postgres,
		redis,
	)
)