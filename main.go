package main

import (
	"log"
	"temporal-robots/workflows"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	c, err := client.Dial(client.Options{
		HostPort: client.DefaultHostPort,
	})
	if err != nil {
		log.Fatalln("Unable to create Temporal client", err)
	}
	defer c.Close()

	w := worker.New(c, "renew-queue", worker.Options{
		// WorkerActivitiesPerSecond:          config.WorkerActivitiesPerSecond,
		// MaxConcurrentActivityExecutionSize: config.MaxConcurrentActivityExecutionSize,
		// Interceptors: []internal.WorkerInterceptor{...},
	})

	w.RegisterWorkflow(workflows.WorkflowDefinition)
	w.RegisterActivity(workflows.GetDomainStatus)
	w.RegisterActivity(workflows.CheckPremiumStatus)
	...

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start Worker", err)
	}
}
