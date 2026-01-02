package main

import (
	"fmt"
	"sync"

	config "singleton.example/internal"
)

func main() {
	var wg sync.WaitGroup
	for i := range 10 {
		wg.Go(func() {
			cfg := config.GetConfig()
			fmt.Printf("From %d -> AppName: %v, Version: %v, Address: %p\n", i, cfg.AppName, cfg.Version, cfg)
		})
	}

	wg.Wait()
}
