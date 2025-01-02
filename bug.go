```go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup
        var counter int
        const numRoutines = 1000

        for i := 0; i < numRoutines; i++ {
                wg.Add(1)
                go func() {
                        defer wg.Done()
                        counter++
                }()
        }

        wg.Wait()
        fmt.Printf("Counter value: %d\n", counter)
}
```