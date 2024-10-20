package main

import (
	"fmt"
	"math/rand"
)

func boring(msg string, quit chan string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; true; i++ {
			select {
			case c <- fmt.Sprintf("%s %d", msg, i):
				// Do nothing
			case <-quit:
				quit <- "See you!"
				return
			}
		}
	}()
	return c
}

func main() {
	quit := make(chan string)
	c := boring("joe", quit)
	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-c)
	}
	quit <- "Bye!"
	fmt.Printf("Joe says: %q\n", <-quit) // This blocks main until the goroutine confirms it's done
}
