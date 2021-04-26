package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		i := 0
		defer wg.Done()
		for {

			if i == 10 {
				return
			}
			fmt.Println("HELLO")

			i += 1
		}
	}(&wg)
	println("waiting")
	wg.Wait()
	println("done")
}
