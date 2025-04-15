package service

import (
	"sync"
)

// Task 表示一个需要处理的任务
type Task func() error

// WorkerPool 管理一组工作goroutine
type WorkerPool struct {
	tasks   chan Task
	wg      sync.WaitGroup
	workers int
}

// NewWorkerPool 创建新的工作池
func NewWorkerPool(workers int) *WorkerPool {
	return &WorkerPool{
		tasks:   make(chan Task, workers*2),
		workers: workers,
	}
}

// Start 启动工作池
func (p *WorkerPool) Start() {
	for i := 0; i < p.workers; i++ {
		p.wg.Add(1)
		go func() {
			defer p.wg.Done()
			for task := range p.tasks {
				task()
			}
		}()
	}
}

// Submit 提交任务到工作池
func (p *WorkerPool) Submit(task Task) {
	p.tasks <- task
}

// Stop 停止工作池
func (p *WorkerPool) Stop() {
	close(p.tasks)
	p.wg.Wait()
}
