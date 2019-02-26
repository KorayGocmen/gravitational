package main

import (
	"sync"

	"github.com/google/uuid"
)

// Currently the workers are kept in an in-memory model.
// In a real production system there would be a need for a
// persistent way to keep information about the online workers.
// Such a persistent model would require two key features:
// 		1. Keeping the data in a persistent database and loading this
// 			 data to memory in the event of a scheduler crash/restart etc.
// 		2. Continuous monitoring of the workers in the cluster by
// 			the scheduler in order to be informed by a potential worker
// 			crash. When workers are exiting normally, they deregister,
// 			however in the event of a crash, the scheduler would be
// 			unaware of the worker status. A heartbeat/ping monitor between
// 			the scheduler and online workers would address this issue.

// workers holds the pointers to all the registered workers.
// 		- key: workerID
// 		- val: pointer to worker object.
var (
	workersMutex = &sync.Mutex{}
	workers      = make(map[string]*worker)
)

// worker holds the information about registered workers
// 		- id: uuid assigned when the worker first register.
// 		- addr: workers network address, later used to create
// 				grpc client to the worker.
type worker struct {
	id   string
	addr string
}

// newWorker creates a new worker instance and adds
// the new worker to the map.
// Returns:
// 		- string: worker id
func newWorker(address string) string {
	workersMutex.Lock()
	defer workersMutex.Unlock()

	workerID := uuid.New().String()
	workers[workerID] = &worker{
		id:   workerID,
		addr: address,
	}

	return workerID
}

// delWorker removes the worker with the given id
// from known workers map.
func delWorker(id string) {
	workersMutex.Lock()
	defer workersMutex.Unlock()

	if _, ok := workers[id]; ok {
		delete(workers, id)
	}
}
