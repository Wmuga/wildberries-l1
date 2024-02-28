package main

import (
	"context"
	"fmt"
	"time"
)

/*
 * Реализовать собственную функцию sleep
 */

func sleep(time time.Duration) {
	// Используем таймаут конекста, чтобы знать когда "просыпаться"
	ctx, cancel := context.WithTimeout(context.Background(), time)
	defer cancel()
	<-ctx.Done()
}

func main() {
	fmt.Println("1")
	sleep(time.Second * 5)
	fmt.Println("2")
}
