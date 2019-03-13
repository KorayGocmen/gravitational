package main

import (
	"context"
	"log"
	"net"

	pb "github.com/koraygocmen/gravitational/jobscheduler"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// server holds the GRPC worker server instance.
type server struct{}

// StartJob starts a new job with the given command and the path
// Command can be any exectuable command on the worker and the path
// is the relative path of the script.
func (s *server) StartJob(ctx context.Context, r *pb.StartJobReq) (*pb.StartJobRes, error) {
	jobID, err := startScript(r.Command, r.Path)
	if err != nil {
		return nil, err
	}

	res := pb.StartJobRes{
		JobID: jobID,
	}

	return &res, nil
}

// StopJob stops a running job with the given job id.
func (s *server) StopJob(ctx context.Context, r *pb.StopJobReq) (*pb.StopJobRes, error) {
	if err := stopScript(r.JobID); err != nil {
		return nil, err
	}

	return &pb.StopJobRes{}, nil
}

// QueryJob returns the status of job with the given job id.
// The status of the job is inside the `Done` variable in response
// and it specifies if the job is still running (true), or stopped (false).
func (s *server) QueryJob(ctx context.Context, r *pb.QueryJobReq) (*pb.QueryJobRes, error) {
	jobDone, jobError, jobErrorText, err := queryScript(r.JobID)
	if err != nil {
		return nil, err
	}

	res := pb.QueryJobRes{
		Done:      jobDone,
		Error:     jobError,
		ErrorText: jobErrorText,
	}
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
	pb.RegisterWorkerServer(grpcServer, &server{})
	grpcServer.Serve(lis)
}
