package main

import (
	"math/rand"
	"time"
)

func boring(msg string) {
	for i := 0; ; i++ {
		println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func main() {
	boring("boring!")
}
