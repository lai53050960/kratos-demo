// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"go.opentelemetry.io/otel/trace"
	"kratos-client/internal/biz"
	"kratos-client/internal/conf"
	"kratos-client/internal/data"
	"kratos-client/internal/server"
	"kratos-client/internal/service"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, log.Logger, trace.TracerProvider) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
