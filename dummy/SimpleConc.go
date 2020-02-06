package main

import (
	"fmt"
	"sync"
)

func main() {
	counter := 0

	c := sync.Mutex{}
	for {
		go func() {
			c.Lock()
			defer c.Unlock()

			counter = counter + 1
			fmt.Println(counter)


		}()
	}

}