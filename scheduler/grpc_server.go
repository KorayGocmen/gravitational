package main

import (
	"context"
	"crypto/subtle"
	"errors"
	"log"
	"net"

	pb "github.com/koraygocmen/gravitational/jobscheduler"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// server holds the GRPC scheduler server instance.
type server struct{}

// RegisterWorker registers a new worker on the server.
// Workers call this method when they are coming online.
func (s *server) RegisterWorker(ctx context.Context, r *pb.RegisterReq) (*pb.RegisterRes, error) {

	// In a production system simple string comparison of an api key
	// provided in run time is not the best way to handle authorization.
	// Shared api keys across different applications create many challenges
	// from a security and maintanance standpoint. A system like Vault to
	// distribute the api key would be needed on a prod system.
	if subtle.ConstantTimeCompare([]byte(r.ApiKey), []byte(apiKey)) == 0 {
		return nil, errors.New("api key unauthorized")
	}

	workerID := newWorker(r.Address)

	res := pb.RegisterRes{
		WorkerID: workerID,
	}

	log.Printf("New worker with ID: %s is online\n", workerID)
	return &res, nil
}

// DeregisterWorker deregisters an existing worker on the server.
// Workers call this method when they are going offline.
func (s *server) DeregisterWorker(ctx context.Context, r *pb.DeregisterReq) (*pb.DeregisterRes, error) {

	delWorker(r.WorkerID)

	res := pb.DeregisterRes{
		Success: true,
	}

	log.Printf("Worker with ID: %s is offline\n", r.WorkerID)
	return &res, nil
}

// startGRPCServer starts a scheduler server instance on the address specified
// by the grpcServerAddr, if the grpcServerUseTLS is true, the
// GRPC server will start with TLS with the key and crt file speficied.
func startGRPCServer() {
	lis, err := net.Listen("tcp", grpcServerAddr)
	if err != nil {
		errs <- err
		return
	}

	// Start with TLS if option is set.
	var opts []grpc.ServerOption
	if grpcServerUseTLS {
		creds, err := credentials.NewServerTLSFromFile(
			grpcServerCrtFile,
			grpcServerKeyFile,
		)
		if err != nil {
			errs <- err
			return
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}

	log.Println("GRPC Server listening on", grpcServerAddr)

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterSchedulerServer(grpcServer, &server{})
	grpcServer.Serve(lis)
}
