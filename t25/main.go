package main

import (
	"context"
	"fmt"
	"time"
)

func sleep(time time.Duration) {
	// Используем таймаут конекста, чтобы знать когда "возвращаться к работе"
	ctx, cancel := context.WithTimeout(context.Background(), time)
	defer cancel()
	<-ctx.Done()
}

func main() {
	fmt.Println("1")
	sleep(time.Second * 5)
	fmt.Println("2")
}
