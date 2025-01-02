```go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup
        var counter int
        var mu sync.Mutex // Add mutex for synchronization
        const numRoutines = 1000

        for i := 0; i < numRoutines; i++ {
                wg.Add(1)
                go func() {
                        defer wg.Done()
                        mu.Lock() // Lock the mutex before accessing the counter
                        counter++
                        mu.Unlock() // Unlock the mutex after accessing the counter
                }()
        }

        wg.Wait()
        fmt.Printf("Counter value: %d\n", counter)
}
```