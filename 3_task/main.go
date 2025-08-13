package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s <workers>\n", os.Args[0])
		os.Exit(1)
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n <= 0 {
		fmt.Fprintln(os.Stderr, "workers must be a positive integer")
		os.Exit(1)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	jobs := make(chan int, n)
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 1; i <= n; i++ {
		go worker(ctx, i, jobs, &wg)
	}
	go func() {
		defer close(jobs)
		val := 1
		ticker := time.NewTicker(200 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				jobs <- val
				val++
			}
		}
	}()
	wg.Wait()
}

func worker(ctx context.Context, id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case job, ok := <-jobs:
			if !ok {
				return
			}
			// Имитация работы
			fmt.Printf("#%d is working on %d\n", id, job)
		}
	}
}
