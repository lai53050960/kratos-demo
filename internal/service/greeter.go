package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel/sdk/trace"
	v1 "kratos-client/api/helloworld/v1"
	"kratos-client/internal/biz"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer

	uc     *biz.GreeterUsecase
	log    *log.Helper
	tracer *trace.TracerProvider
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase, logger log.Logger, provider *trace.TracerProvider) *GreeterService {
	return &GreeterService{uc: uc, log: log.NewHelper(logger), tracer: provider}
}

// SayHello implements helloworld.GreeterServer
func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	//foo := errors.BadRequest(v1.ErrorReason_ERROR_REASON_UNSPECIFIED.String(), "aaa")
	//fmt.Println(foo)
	//conn, err := grpc.DialInsecure(ctx,
	//	grpc.WithEndpoint("localhost:9502"),
	//	grpc.WithMiddleware(middleware.Chain(
	//		tracing.Client(
	//			tracing.WithTracerProvider(s.tracer),
	//			tracing.WithPropagators(
	//				propagation.NewCompositeTextMapPropagator(propagation.Baggage{}, propagation.TraceContext{}),
	//			),
	//		),
	//		recovery.Recovery())),
	//	grpc.WithTimeout(2*time.Second),
	//)
	//if err != nil {
	//	return nil, err
	//}
	//c := pb.NewUserClient(conn)

	greeter := biz.Greeter{
		Hello: "lai",
	}
	r, err := s.uc.SayHello(ctx, &greeter)
	if err != nil {
		s.log.Infof("could not greet: %v", err)
	}
	return &v1.HelloReply{Message: "Hello " + r.Name}, nil
}
