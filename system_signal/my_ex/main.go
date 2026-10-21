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

func Worker(ctx context.Context, n int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Рабочий под номером", n, "закончил работу")
			return
		case <-time.After(1 * time.Second):
			fmt.Println("Рабочий под номером", n, "работает")
		}
	}
}

func main() {
	fmt.Println("PID:", os.Getpid())

	workers_count, err := strconv.Atoi(os.Getenv("WORKERS_COUNT"))

	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGTERM)

	if err != nil {
		fmt.Println("fail to convert to integer:", err.Error())
		return
	}

	/* go func() {
		time.Sleep(5 * time.Second)
		ctxCancel()
		fmt.Println("Рабочий день закончен")
	}() */

	wg := &sync.WaitGroup{}

	for i := 1; i <= workers_count; i++ {
		wg.Add(1)
		go Worker(ctx, i, wg)
	}

	wg.Wait()
}
