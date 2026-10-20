package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("PID:", os.Getpid())

	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGTERM)

	fmt.Println("До отмены контекста")
	<-ctx.Done()
	fmt.Println("После отмены контекста")

}
