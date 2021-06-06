package hw05parallelexecution //nolint:golint,stylecheck

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	errLimit := m
	taskCh := make(chan Task, len(tasks))
	wg := sync.WaitGroup{}
	mx := sync.Mutex{}
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			worker(taskCh, &mx, &errLimit)
		}()
	}

	for _, t := range tasks {
		taskCh <- t
	}

	close(taskCh)
	wg.Wait()

	if errLimit < 0 {
		return ErrErrorsLimitExceeded
	}

	return nil
}

func worker(taskCh <-chan Task, mx sync.Locker, errLimit *int) {
	stop := false
	for t := range taskCh {
		err := t()
		mx.Lock()
		if *errLimit < 1 {
			stop = true
		}
		if err != nil {
			*errLimit--
		}
		mx.Unlock()
		if stop {
			return
		}
	}
}
