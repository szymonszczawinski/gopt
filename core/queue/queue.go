package queue

import (
	"context"
	"log"
	// "sync"

	"golang.org/x/sync/errgroup"
)

type Job struct {
	Execute func()
}

type IJobQueue interface {
	Wait()
}

type JobQueue struct {
	// wg   sync.WaitGroup
	jobs chan *Job
}

func NeqJobQueue() *JobQueue {
	log.Println("NewJobQueue")
	return &JobQueue{jobs: make(chan *Job)}
}

func (jq *JobQueue) Wait() {
	log.Println("JobQueue::Wait")
	// jq.wg.Wait()
}

// Start starts a dispatcher.
// This dispatcher will stops when it receive a value from `ctx.Done`.
func (jq *JobQueue) Start(g *errgroup.Group, ctx context.Context) {
	log.Println("JobQueue::Start")
	g.Go(func() error {
	Loop:
		for {
			log.Println("JobQueue::Start::Wait for Job")
			select {
			case <-ctx.Done():
				log.Println("JobQueue::Start::Finish")
				// wg.Done()
				break Loop

			case job := <-jq.jobs:

				log.Println("JobQueue::Start::Do Job")
				job.Execute()

			}
		}
		log.Println("Done")
		return nil
	})
}

// Add enqueues a job into the queue.
// If the number of enqueued jobs has already reached to the maximum size,
// this will block until the other job has finish and the queue has space to accept a new job.
func (jq *JobQueue) Add(job *Job) {
	log.Println("JobQueue::Add")
	jq.jobs <- job
}

func (jq *JobQueue) Stop() {
	// jq.wg.Done()
}

func (jq *JobQueue) loop(ctx context.Context) {
	// wg := sync.WaitGroup{}
	// wg.Add(1)
	// go func() {
	// }()
	// log.Println("JobQueue::loop::Wait")
	// wg.Wait()
	log.Println("End")
	// jq.Stop()
}
