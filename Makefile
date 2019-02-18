# Makefile

SCHEDULER_PATH = scheduler
SCHEDULER_BINARY = scheduler

WORKER_PATH = worker
WORKER_BINARY = worker

PROTO_PATH = jobscheduler
PROTO_NAME = job_scheduler.proto
PROTO_PLUGIN = jobscheduler

.PHONY: all
all: clean build

.PHONY: build
build: build_scheduler build_worker

.PHONY: build_proto
build_proto:
	protoc -I $(PROTO_PATH)/ $(PROTO_PATH)/$(PROTO_NAME) --go_out=plugins=grpc:$(PROTO_PLUGIN)

.PHONY: build_scheduler
build_scheduler:
	go build -o $(SCHEDULER_PATH)/$(SCHEDULER_BINARY) $(SCHEDULER_PATH)/*.go

.PHONY: build_worker
build_worker:
	go build -o $(WORKER_PATH)/$(WORKER_BINARY) $(WORKER_PATH)/*.go

.PHONY: clean
clean: 
	rm -f *.out
	rm -f $(SCHEDULER_PATH)/$(SCHEDULER_BINARY)
	rm -f $(WORKER_PATH)/$(WORKER_BINARY)
