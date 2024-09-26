package calculator_test

import (
	"context"
	"testing"

	calculator "grpc-calculator/internal/grpc/calculator"

	"github.com/stretchr/testify/assert"
	pb "github.com/vshulcz/grpc-protos/gen/go/calculator"
)

func TestAdd(t *testing.T) {
	server := &calculator.ServerAPI{}

	req := &pb.NumbersRequest{
		Number1: 5,
		Number2: 3,
	}

	resp, err := server.Add(context.Background(), req)

	assert.NoError(t, err)
	assert.Equal(t, float32(8), resp.Result)
}

func TestSubtract(t *testing.T) {
	server := &calculator.ServerAPI{}

	req := &pb.NumbersRequest{
		Number1: 10,
		Number2: 4,
	}

	resp, err := server.Subtract(context.Background(), req)

	assert.NoError(t, err)
	assert.Equal(t, float32(6), resp.Result)
}

func TestMultiply(t *testing.T) {
	server := &calculator.ServerAPI{}

	req := &pb.NumbersRequest{
		Number1: 2,
		Number2: 7,
	}

	resp, err := server.Multiply(context.Background(), req)

	assert.NoError(t, err)
	assert.Equal(t, float32(14), resp.Result)
}

func TestDivide(t *testing.T) {
	server := &calculator.ServerAPI{}

	req := &pb.NumbersRequest{
		Number1: 10,
		Number2: 2,
	}

	resp, err := server.Divide(context.Background(), req)

	assert.NoError(t, err)
	assert.Equal(t, float32(5), resp.Result)
}

func TestDivideByZero(t *testing.T) {
	server := &calculator.ServerAPI{}

	req := &pb.NumbersRequest{
		Number1: 10,
		Number2: 0,
	}

	_, err := server.Divide(context.Background(), req)
	assert.Error(t, err)
}
