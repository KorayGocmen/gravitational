package main

import (
	"context"
	"log"

	pb "github.com/koraygocmen/gravitational/jobscheduler"
	"google.golang.org/grpc"
)

// registerWorker dials the scheduler GRPC server and registers
// the calling worker with the worker's GRPC server address.
// Worker's GRPC server address is later used by the scheduler to dial
// worker to start/stop/query jobs.
func registerWorker(ctx context.Context) {
	// TODO: Change to secure when basic functionality is done.
	conn, err := grpc.Dial(schedulerAddr, grpc.WithInsecure())
	if err != nil {
		errs <- err
		return
	}
	defer conn.Close()
	c := pb.NewSchedulerClient(conn)

	registerReq := pb.RegisterReq{
		ApiKey:  apiKey,
		Address: grpcServerAddr,
	}
	r, err := c.RegisterWorker(ctx, &registerReq)
	if err != nil {
		errs <- err
		return
	}

	workerID = r.WorkerID
	log.Printf("Registered on scheduler with ID: %s", r.WorkerID)
}

// deregisterWorker deregisters the calling worker from the scheduler.
// Scheduler will remove the worker from the known workers. Any nonpanic
// exit by the worker application should be calling deregister function
// before termination.
func deregisterWorker(ctx context.Context) {
	// TODO: Change to secure when basic functionality is done.
	conn, err := grpc.Dial(schedulerAddr, grpc.WithInsecure())
	if err != nil {
		// deregisterWorker errors can not be sent to the errs
		// channel since the errs channel will be calling
		// deregisterWorker again, causing an infinite loop.
		log.Println(err)
		return
	}
	defer conn.Close()
	c := pb.NewSchedulerClient(conn)

	deregisterReq := pb.DeregisterReq{
		WorkerID: workerID,
	}
	r, err := c.DeregisterWorker(ctx, &deregisterReq)
	if err != nil {
		// deregisterWorker errors can not be sent to the errs
		// channel since the errs channel will be calling
		// deregisterWorker again, causing an infinite loop.
		log.Println(err)
		return
	}

	log.Printf("Deregistered OK: %t", r.Success)
}
