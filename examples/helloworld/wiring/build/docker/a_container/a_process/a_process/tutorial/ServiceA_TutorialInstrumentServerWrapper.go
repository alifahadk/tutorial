// Blueprint: Auto-generated by Tutorial Plugin
package tutorial

import (
	"context"
	"errors"
	"fmt"
	"github.com/blueprint-uservices/tutorial/examples/helloworld/workflow/servicea"
	"log"
)

type Job struct {
	Method  func(context.Context) error
	Context context.Context
	Done    chan struct{}
}

func StartWorkerPool(workerCount, queueSize int) chan Job {
	jobQueue := make(chan Job, queueSize)

	for i := 1; i <= workerCount; i++ {
		go func(id int, jobs <-chan Job) { //Read-only channel using <-chan Job instead of chan Job
			for job := range jobs {
				err := job.Method(job.Context)
				if err != nil {
					fmt.Println("Error when executing job:", err)
				} else {
					fmt.Println("Job executed successfully")
				}
				close(job.Done)
			}
		}(i, jobQueue)
	}

	return jobQueue
}

type ServiceA_TutorialInstrumentServerWrapper struct {
	Service servicea.ServiceA
	Queue   chan Job
}

func New_ServiceA_TutorialInstrumentServerWrapper(ctx context.Context, service servicea.ServiceA) (*ServiceA_TutorialInstrumentServerWrapper, error) {
	handler := &ServiceA_TutorialInstrumentServerWrapper{}
	handler.Service = service
	handler.Queue = StartWorkerPool(4, 10)
	return handler, nil
}

func (handler *ServiceA_TutorialInstrumentServerWrapper) Hello(ctx context.Context) error {
	log.Println("Processing Hello")

	done := make(chan struct{})
	job := Job{Method: handler.Service.Hello, Context: ctx, Done: done}

	select {
	case handler.Queue <- job:
		<-done // Wait for job to complete
		return nil
	default:
		return errors.New("Job not executed successfully")
	}
}

func (handler *ServiceA_TutorialInstrumentServerWrapper) World(ctx context.Context) error {
	log.Println("Processing World")

	done := make(chan struct{})
	job := Job{Method: handler.Service.World, Context: ctx, Done: done}

	select {
	case handler.Queue <- job:
		<-done // Wait for job to complete
		return nil
	default:
		return errors.New("Job not executed successfully")
	}
	return handler.Service.World(ctx)
}
