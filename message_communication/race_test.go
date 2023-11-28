package main

import (
	_"sync"
	"sync/atomic"
	"testing"
)

// func TestDataRaceCondition(t *testing.T) {
// 	var state int32
// 	var mu sync.RWMutex

// 	//For avoiding race conditions use Mutexes or Lock
// 	for i := 0; i < 10; i++ {
// 		go func(i int) {
// 			mu.Lock()  //Here we are using Read-Write Mutexes and here we are spinning up 10 goroutines.
// 			// Basically the goroutines that holds the lock can read and write. And here the lock is of global type.
// 			state += 32
// 			//This is after the business logic
// 			mu.Unlock()
// 		}(i)
// 	}
// }

// One thing to keep in mind is Mutexes slow down the performance of the system. Alternatively use the channels wisely.
// Or else use the atomic values.

func TestDataRaceCondition(t *testing.T) {
	var state int32

	for i:=0; i < 10; i++ {
		go func(i int) {
			atomic.AddInt32(&state, int32(i))
		}(i)
	}
}