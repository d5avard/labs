### Summary

This Go program demonstrates the memory consumption of goroutines by measuring the memory used before and after launching a large number of goroutines. Here's a summary from that perspective:

1. Memory Measurement Function:
    - memConsumed: A function that triggers garbage collection, reads memory statistics, and returns the total system memory used.

2. Channel and WaitGroup:
    - c: A read-only channel of empty interfaces.
    - wg: A WaitGroup to synchronize the completion of goroutines.

3. No-Operation Goroutine:
    - noop: A function that marks the goroutine as done in the WaitGroup and then waits indefinitely on the channel c.

4. Number of Goroutines:
    - numGoroutines: A constant defining the number of goroutines to launch (10,000 in this case).

5. Memory Consumption Before:
    - before: Memory consumed before launching the goroutines.

6. Launching Goroutines:
    - A loop that launches numGoroutines goroutines, each running the noop function.

7. Waiting for Goroutines:
    - wg.Wait(): Waits for all launched goroutines to call wg.Done().

8. Memory Consumption After:
    - after: Memory consumed after all goroutines have been launched.


9. Memory Consumption Per Goroutine:
    - Prints the average memory consumption per goroutine in kilobytes.

This program calculates the average memory consumption of a single goroutine by launching 10,000 goroutines, measuring the memory used before and after, and then computing the difference. The result is printed in kilobytes.