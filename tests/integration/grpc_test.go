package calculator_test

import (
	"context"
	"net"
	"testing"

	"grpc-calculator/internal/grpc/calculator"

	"github.com/stretchr/testify/assert"
	pb "github.com/vshulcz/grpc-protos/gen/go/calculator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGRPCServer(t *testing.T) {
	listener, err := net.Listen("tcp", ":0")
	assert.NoError(t, err)

	server := grpc.NewServer()
	calculator.Register(server)

	errCh := make(chan error, 1)
	go func() {
		if err := server.Serve(listener); err != nil {
			errCh <- err
		}
		close(errCh)
	}()
	defer server.Stop()

	select {
	case err := <-errCh:
		if err != nil {
			t.Fatalf("Failed to serve: %v", err)
		}
	default:
	}

	conn, err := grpc.NewClient(listener.Addr().String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	assert.NoError(t, err)
	defer conn.Close()

	client := pb.NewCalculatorClient(conn)

	t.Run("Add", func(t *testing.T) {
		resp, err := client.Add(context.Background(), &pb.NumbersRequest{
			Number1: 5,
			Number2: 3,
		})
		assert.NoError(t, err)
		assert.Equal(t, float32(8), resp.Result)
	})

	t.Run("Divide", func(t *testing.T) {
		resp, err := client.Divide(context.Background(), &pb.NumbersRequest{
			Number1: 10,
			Number2: 2,
		})
		assert.NoError(t, err)
		assert.Equal(t, float32(5), resp.Result)
	})

	t.Run("DivideByZero", func(t *testing.T) {
		_, err := client.Divide(context.Background(), &pb.NumbersRequest{
			Number1: 10,
			Number2: 0,
		})
		assert.Error(t, err)
	})
}
