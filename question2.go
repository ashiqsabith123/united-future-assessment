package main

import "sync"

var (
	mutex sync.RWMutex
	name  string
)

func question2() {
	mutex.Lock()
	//defer works after the function execution
	defer mutex.Unlock()
	// we can safely perform the read or write operation here
	name = "something"
}
