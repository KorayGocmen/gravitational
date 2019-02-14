# Scheduler-Worker Cluster Architecture


# Goal

*   Build a scheduler and worker architecture that supports GRPC communication.

# General Overview

*   Support single scheduler and multiple workers.
*   The scheduler must be aware of all workers' states. (requires the workers to register and deregister when going offline)
*   The scheduler and the workers authenticate with each other via SSL.
*   The scheduler and workers, both need to have GRPC-APIs. 
    *   The scheduler GRPC-API support:
        *   Register a worker
        *   Deregister a worker
    *   The worker GRPC-API support:
        *   Start a job
        *   Stop a job
        *   Return the status of a job
        *   Return a stream of output for a running job.


# Scheduler Overview

* [Scheduler Details](scheduler/scheduler.md)
* All config parameters are specified in the config.toml file
* When a worker registers, a UUID is assigned to the worker and worker details are kept in a map.
* Starting a job on a specific worker
  * Request: 
    ```
    {
      "worker_id": "71382ed1-471d-4ae3-b572-f67d178f04e9",
      "path": "worker/scripts/count.sh",
      "command": "bash"
    }
    ```
  * Response: 
    ```
    {
      "job_id": "6c26a00e-3017-11e9-b210-d663bd873d93"
    }
    ```


# Worker Overview

* [Worker Details](worker/worker.md)
* All config parameters are specified in the config.toml file
* Support a GRPC-API (Appendix B)
* Jobs are basically scripts that are held in the specified folder in the config.
* When a job is started by the scheduler, a job object is created by the worker
  * This object specifies where the output of the job will be piped.
  * Also holds which command and which path was requested.