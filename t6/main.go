package main

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

/*
 * Реализовать все возможные способы остановки выполнения горутины.
 */

// Воркер с закрытием по контексту
func wrkCtx(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("Ctx worker workds")
			time.Sleep(time.Second / 3)
		}
	}
}

// Воркер с закрытием по каналу
func wrkChan(c chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-c: // можно еще _, ok := <- c if !ok{}
			return
		default:
			fmt.Println("Chan worker works")
			time.Sleep(time.Second / 3)
		}
	}
}

// Воркер с закрытием через runtime
func wrkRuntime(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		fmt.Println("Runtime worker works")
		time.Sleep(time.Second / 3)
		// Когда-нибудь закроется
		if rand.Intn(3) == 0 {
			runtime.Goexit()
		}
	}
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(3)
	c := make(chan struct{})
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	go wrkChan(c, wg)
	go wrkCtx(ctx, wg)
	go wrkRuntime(wg)
	<-ctx.Done()
	c <- struct{}{}
	wg.Wait()
	fmt.Println("Done")
}
