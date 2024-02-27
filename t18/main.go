package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mux     *sync.RWMutex
	counter int
}

func NewCounter() *Counter {
	return &Counter{
		mux: &sync.RWMutex{},
	}
}

func (c *Counter) Add(num int) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.counter += num
}

func (c *Counter) Value() int {
	c.mux.RLock()
	defer c.mux.RUnlock()
	return c.counter
}

// Будет увеличивать счетчик на 1 {count} раз
func worker(c *Counter, count int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < count; i++ {
		c.Add(1)
	}
}

func main() {
	wg := &sync.WaitGroup{}
	c := NewCounter()
	// Запуск 20 воркеров, которые увеличат счетчик 1000 раз
	wg.Add(20)
	for i := 0; i < 20; i++ {
		go worker(c, 1000, wg)
	}
	wg.Wait()
	fmt.Println(c.Value())
}
