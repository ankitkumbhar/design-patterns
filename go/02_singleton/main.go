package main

import (
	"fmt"
	"sync"
	"time"
)

var obj *DriverPg
var lock = sync.Mutex{}

type DriverPg struct {
	Conn string
}

func connect(str string) *DriverPg {
	// thread safe handling for creating only single instance
	// sync lock will lock below block so that only one goroutine will access it
	lock.Lock()

	defer lock.Unlock()

	if obj == nil {
		fmt.Println("Creating object with goroutine no : ", str)
		obj = &DriverPg{
			Conn: fmt.Sprintf("connection-string :"),
		}
	}

	return obj
}

func main() {
	// creating object using go routines
	for i := 0; i < 100; i++ {
		go func(idx int) {
			go connect(fmt.Sprintf("goroutine : %v", idx))
		}(i)
	}

	time.Sleep(time.Second * 2)
}
