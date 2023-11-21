package queue

import (
	"context"
	"log"

	"golang.org/x/sync/errgroup"
)

type Job struct {
	Execute func()
}

type IJobQueue interface {
	Start(context.Context)
	Add(*Job)
}

type jobQueue struct {
	name string
	g    *errgroup.Group
	jobs chan *Job
}

func NeqJobQueue(name string, g *errgroup.Group) IJobQueue {
	log.Println("NewJobQueue::", name)
	return &jobQueue{
		jobs: make(chan *Job),
		g:    g,
		name: name,
	}
}

// Start starts a dispatcher.
// This dispatcher will stops when it receive a value from `ctx.Done`.
func (jq *jobQueue) Start(ctx context.Context) {
	log.Println("JobQueue", jq.name, "::Start")
	jq.g.Go(func() error {
		defer close(jq.jobs)
	Loop:
		for {
			log.Println("JobQueue", jq.name, "::Wait for Job")
			select {
			case <-ctx.Done():
				log.Println("JobQueue", jq.name, "::Finish")
				break Loop

			case job := <-jq.jobs:

				log.Println("JobQueue", jq.name, "::Do Job")
				jq.g.Go(func() error {
					job.Execute()
					return nil
				})

			}
		}
		log.Println("JobQueue", jq.name, "::Done")
		return nil
	})
}

// Add enqueues a job into the queue.
// If the number of enqueued jobs has already reached to the maximum size,
// this will block until the other job has finish and the queue has space to accept a new job.
func (jq *jobQueue) Add(job *Job) {
	log.Println("JobQueue", jq.name, "::Add")
	jq.g.Go(func() error {

		jq.jobs <- job
		return nil
	})
	log.Println("JobQueue", jq.name, "::Added")
}
