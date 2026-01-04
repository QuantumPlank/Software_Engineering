package main

import (
	"fmt"

	"proxy.example/internal/user"
)

func main() {
	realService := user.NewRealService()
	proxy := user.NewServiceProxy(realService)

	fmt.Println(proxy.GetUser(10))
	fmt.Println(proxy.GetUser(1))
	fmt.Println(proxy.GetUser(10))
	fmt.Println(proxy.GetUser(1))
}
