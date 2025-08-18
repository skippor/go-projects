package main

import (
    "context"
    "fmt"
    "sync"
    "time"
    "runtime"
)

const (
    _maxProc = 4
    _tickerInterval = 1
)

func worker(ctx context.Context, wg *sync.WaitGroup) {
    defer wg.Done()

    ticker := time.NewTicker(_tickerInterval * time.Second)
    defer ticker.Stop()

    for {
        select {
        case <-ctx.Done():
            fmt.Println("goroutine finished")
            return
        case <-ticker.C:
            fmt.Println("goroutine ticker done")
        default:
            fmt.Println("goroutine running")
            time.Sleep(500 * time.Millisecond)
        }
    }
}

func main() {
    runtime.GOMAXPROCS(_maxProc)

    ctx, cancel := context.WithCancel(context.Background())
    wg := &sync.WaitGroup{}

    wg.Add(1)
    go worker(ctx, wg)

    time.Sleep(5 * time.Second)
    cancel()
    
    wg.Wait()
}