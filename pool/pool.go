package pool

import (
	"sync"
)

type Task interface {
	Execute()
}

/**
 * Structure of workers pool
 */
type Pool struct {
	size  int
	tasks chan Task
	wg    sync.WaitGroup
}

/**
 * Return a pool of workers, according to the mentioned size
 */
func NewPool(size int) *Pool {
	pool := &Pool{
		tasks: make(chan Task, 512),
	}
	pool.create(size)
	return pool
}

/**
 * Create pool of workers, according to the mentioned size
 */
func (p *Pool) create(size int) {
	for p.size < size {
		p.size++
		p.wg.Add(1)
		go p.worker()
	}
}

/**
 * Creates pool worker
 */
func (p *Pool) worker() {
	defer p.wg.Done()
	for {
		task, ok := <-p.tasks
		if !ok {
			return
		}
		task.Execute()
	}
}

/**
 * Close pool
 */
func (p *Pool) Close() {
	close(p.tasks)
}

/**
 * Wait for pool workers to complete execution
 */
func (p *Pool) Wait() {
	p.wg.Wait()
}

/**
 * Execute a particular task
 */
func (p *Pool) Exec(task Task) {
	p.tasks <- task
}

/**
 * Get current queue size
 */
func (p *Pool) GetQueueSize() int {
	return len(p.tasks)
}
