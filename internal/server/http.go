package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/trace"
	v1 "kratos-client/api/helloworld/v1"
	"kratos-client/api/user"
	"kratos-client/internal/conf"
	"kratos-client/internal/service"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, greeter *service.GreeterService, userService *service.UserService, logger log.Logger, tracer *trace.TracerProvider) *http.Server {
	var opts = []http.ServerOption{}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}

	m := http.Middleware(
		middleware.Chain(
			recovery.Recovery(recovery.WithLogger(logger)),
			tracing.Server(
				tracing.WithTracerProvider(tracer),
				tracing.WithPropagator(
					propagation.NewCompositeTextMapPropagator(propagation.Baggage{}, propagation.TraceContext{}),
				),
			),
			logging.Server(logger),
		),
	)

	opts = append(opts, m)
	srv := http.NewServer(opts...)

	v1.RegisterGreeterHTTPServer(srv, greeter)
	user.RegisterUserHTTPServer(srv, userService)

	return srv
}
