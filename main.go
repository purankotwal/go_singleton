package main

import (
	"fmt"

	singleton_service "example.com/golang/singleton/services"
)

func main() {
	var cs singleton_service.CounterSingleton
	cSvc := cs.CounterSingletonService()
	cSvc.Next()
	fmt.Printf("Counter: [%d]\n", cSvc.Count)
	cSvc.Next()
	fmt.Printf("Counter: [%d]\n", cSvc.Count)
	// reinitialize
	cSvc = cs.CounterSingletonService()
	cSvc.Next()
	fmt.Printf("Counter: [%d]\n", cSvc.Count)
	cSvc.Previous()
	fmt.Printf("Counter: [%d]\n", cSvc.Count)
}
