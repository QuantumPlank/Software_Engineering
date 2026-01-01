package main

import (
	"fmt"
	"sync"
)

type Singleton struct {
	data string
}

var instance *Singleton
var once sync.Once

func GetInstance() *Singleton {
	once.Do(func() {
		fmt.Println("Creating Singleton")
		instance = &Singleton{data: "This is a Singleton"}
	})
	return instance
}

func main() {
	var wg sync.WaitGroup

	for i := range 10 {
		wg.Go(func() {
			s := GetInstance()
			fmt.Printf("From %d: %v\n", i, s.data)
		})
	}

	wg.Wait()
}
