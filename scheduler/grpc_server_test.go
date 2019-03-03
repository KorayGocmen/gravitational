package main

import (
	"context"
	"testing"

	pb "github.com/koraygocmen/gravitational/jobscheduler"
)

func TestRegisterWorker(t *testing.T) {
	s := server{}

	workerAddr := "127.0.0.1:53000"

	// Testing the valid request
	validRes, err := s.RegisterWorker(context.Background(), &pb.RegisterReq{
		ApiKey:  apiKey,
		Address: workerAddr,
	})

	if err != nil {
		t.Error("server returned an error for a valid request", err)
	}

	if workerAddr != workers[validRes.WorkerID].addr {
		t.Error("server did not create worker properly")
	}

	// Testing the invalid request
	_, err = s.RegisterWorker(context.Background(), &pb.RegisterReq{
		ApiKey:  "invalid-api-key",
		Address: workerAddr,
	})

	if err == nil {
		t.Error("invalid api key was not detected")
	}
}

func TestDeregisterWorker(t *testing.T) {
	s := server{}

	workerAddr := "127.0.0.1:53000"

	// Testing the valid request
	registerRes, err := s.RegisterWorker(context.Background(), &pb.RegisterReq{
		ApiKey:  apiKey,
		Address: workerAddr,
	})

	if err != nil {
		t.Error("server returned an error for a valid register request", err)
	}

	// Testing the deregister request
	_, err = s.DeregisterWorker(context.Background(), &pb.DeregisterReq{
		WorkerID: registerRes.WorkerID,
	})

	if err != nil {
		t.Error("server returned an error for a valid deregister request", err)
	}

	if _, ok := workers[workerAddr]; ok {
		t.Error("server did not delete worker properly")
	}
}
