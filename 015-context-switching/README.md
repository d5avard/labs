### Summary
This benchmark measures the time required for context switching between goroutines in Go. Specifically, it calculates the time it takes for a sender goroutine to send a message to a receiver goroutine through a channel, repeated b.N times. The WaitGroup ensures that the benchmark only completes after both goroutines have finished their work, providing an accurate measurement of the context switching overhead.

To execute the test on one CPU.
```
$ go test -bench=. -cpu=1 ./main_test.go
```