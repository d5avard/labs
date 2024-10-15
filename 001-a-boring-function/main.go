package main

import "time"

func boring(msg string) {
	for i := 0; ; i++ {
		println(msg, i)
		time.Sleep(time.Second)
	}
}

func main() {
	boring("boring!")
}
