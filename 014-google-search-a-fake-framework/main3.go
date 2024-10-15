package main

import (
	"fmt"
	"time"

	"math/rand"
)

var (
	Web   = fakeSearch("web")
	Image = fakeSearch("image")
	Video = fakeSearch("video")
)

type Result string

type Search func(query string) Result

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

func Google(query string) (results []Result) {
	c := make(chan Result)

	// Spin each of the searches off into their own goroutine and pipe all results back to the channel created above
	go func() { c <- Web(query) }()
	go func() { c <- Image(query) }()
	go func() { c <- Video(query) }()

	// Now we pull each Result from the channel, but timeout if searches take longer than 80ms
	timeout := time.After(80 * time.Millisecond)
	var result Result
	for i := 0; i < 3; i++ {
		select {
		case result = <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("Timed out")
			return
		}
	}

	return
}

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	results := Google("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}
