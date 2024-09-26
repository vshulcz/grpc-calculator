package calculator

import (
	"context"

	"github.com/vshulcz/grpc-protos/gen/go/calculator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ServerAPI struct {
	calculator.UnimplementedCalculatorServer
}

func Register(gRPC *grpc.Server) {
	calculator.RegisterCalculatorServer(gRPC, &ServerAPI{})
}

func (s *ServerAPI) Add(ctx context.Context, req *calculator.NumbersRequest) (*calculator.OperationResponse, error) {
	result := req.Number1 + req.Number2
	return &calculator.OperationResponse{Result: result}, nil
}

func (s *ServerAPI) Subtract(ctx context.Context, req *calculator.NumbersRequest) (*calculator.OperationResponse, error) {
	result := req.Number1 - req.Number2
	return &calculator.OperationResponse{Result: result}, nil
}

func (s *ServerAPI) Multiply(ctx context.Context, req *calculator.NumbersRequest) (*calculator.OperationResponse, error) {
	result := req.Number1 * req.Number2
	return &calculator.OperationResponse{Result: result}, nil
}

func (s *ServerAPI) Divide(ctx context.Context, req *calculator.NumbersRequest) (*calculator.OperationResponse, error) {
	if req.Number2 == 0 {
		return nil, status.Error(codes.InvalidArgument, "cannot divide by zero")
	}
	result := req.Number1 / req.Number2
	return &calculator.OperationResponse{Result: result}, nil
}
