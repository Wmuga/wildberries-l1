package main

import (
	"fmt"
	"sync"

	"golang.org/x/exp/maps"
)

/*
 * Реализовать конкурентную запись данных в map.
 */

// Ну вообще есть sync.Map. Но если мы хотим прям в map, да и с генериками
type SyncMap[keyT comparable, valueT any] struct {
	m   map[keyT]valueT
	mux *sync.RWMutex
}

// Создание объекта
func NewSyncMap[keyT comparable, valueT any]() *SyncMap[keyT, valueT] {
	return &SyncMap[keyT, valueT]{
		m:   map[keyT]valueT{},
		mux: &sync.RWMutex{},
	}
}

// Добавить значение
func (sm *SyncMap[keyT, valueT]) Add(key keyT, value valueT) {
	sm.mux.Lock()
	defer sm.mux.Unlock()
	sm.m[key] = value
}

// Получить значение
func (sm *SyncMap[keyT, valueT]) Get(key keyT) (valueT, bool) {
	sm.mux.RLock()
	defer sm.mux.RUnlock()
	v, ex := sm.m[key]
	return v, ex
}

// Все ключи
func (sm *SyncMap[keyT, valueT]) Keys() []keyT {
	sm.mux.RLock()
	defer sm.mux.RUnlock()
	return maps.Keys(sm.m)
}

// Все значения
func (sm *SyncMap[keyT, valueT]) Values() []valueT {
	sm.mux.RLock()
	defer sm.mux.RUnlock()
	return maps.Values(sm.m)
}

// Запишет 5 ключей начиная с n в мапу
func writer(sm *SyncMap[int, int], n int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		sm.Add(n+i, i)
	}
}

func main() {
	sm := NewSyncMap[int, int]()
	wg := &sync.WaitGroup{}
	// Запуск 5 рабочих на запись в мапу
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go writer(sm, i*10, wg)
	}
	wg.Wait()
	// Вывод всех ключей. Ожидается 25 ключей
	fmt.Println(len(sm.Keys()), sm.Keys())
}
