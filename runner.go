package gracegrpcserverrunner

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

// GRPCServerRunner runs a GRPC server.
type GRPCServerRunner struct {
	Server        *grpc.Server
	ListenAddress string
}

// Run starts a grpc server.
func (r *GRPCServerRunner) Run(ctx context.Context) error {
	go r.handleCtxDone(ctx)

	lis, err := net.Listen("tcp", r.ListenAddress)
	if err != nil {
		return fmt.Errorf("could not listen on tcp address: %w", err)
	}
	err = r.Server.Serve(lis)
	if err == nil {
		return nil
	}
	// ErrServerStopped is returned when we call r.Server.Close() from handleCtxDone
	// so we shouldn't treat it as an error.
	if err == grpc.ErrServerStopped {
		return nil
	}
	return err
}

func (r *GRPCServerRunner) handleCtxDone(ctx context.Context) {
	<-ctx.Done()
	if r.Server != nil {
		r.Server.GracefulStop()
	}
}
