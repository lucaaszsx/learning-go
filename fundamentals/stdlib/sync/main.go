package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    // ----- WaitGroup: wait for a group of goroutines to finish
    // Add(n): increment the counter by n
    // Done(): decrement the counter by 1 (same as Add(-1))
    // Wait(): blocks until the counter reaches 0
    var wg sync.WaitGroup

    for i := 1; i <= 3; i++ {
        wg.Add(1) // increment counter before starting goroutine
        go func(id int) {
            defer wg.Done() // decrement counter when done
            fmt.Printf("Worker %d starting\n", id)
            time.Sleep(time.Duration(id) * 100 * time.Millisecond)
            fmt.Printf("Worker %d done\n", id)
        }(i)
    }

    wg.Wait() // block until all workers are done
    fmt.Println("All workers completed!")

    // ----- Mutex: protect shared data from concurrent access
    // When multiple goroutines access the same variable, you need a Mutex
    // Lock(): acquire the lock (blocks if already locked)
    // Unlock(): release the lock
    var mu sync.Mutex
    counter := 0

    var wg2 sync.WaitGroup
    for i := 0; i < 1000; i++ {
        wg2.Add(1)
        go func() {
            defer wg2.Done()
            mu.Lock()         // acquire lock before accessing shared data
            counter++         // this is now safe to execute
            mu.Unlock()       // release lock after modifying
        }()
    }

    wg2.Wait()
    fmt.Println("Counter (with mutex):", counter) // always 1000

    // ----- RWMutex: reader-writer lock
    // Multiple readers can access simultaneously
    // Only one writer at a time, and no readers during writing
    var rwMu sync.RWMutex
    data := map[string]string{"key": "value"}

    // Reading: use RLock/RUnlock
    readData := func(key string) string {
        rwMu.RLock()         // acquire read lock
        defer rwMu.RUnlock() // release read lock
        return data[key]
    }

    // Writing: use Lock/Unlock
    writeData := func(key, value string) {
        rwMu.Lock()         // acquire write lock
        defer rwMu.Unlock() // release write lock
        data[key] = value
    }

    writeData("name", "Lucas")
    fmt.Println("Read:", readData("name"))

    // ----- Once: execute something exactly once
    // Useful for initialization that must happen only once
    var once sync.Once

    initialize := func() {
        fmt.Println("Initializing... (this runs only once)")
    }

    // Even if called from multiple goroutines, initialize runs once
    for i := 0; i < 5; i++ {
        once.Do(initialize)
    }
    fmt.Println("Program continues...")

    // ----- sync.Map: concurrent-safe map
    // Use when: keys are mostly stable, or different goroutines read/write different keys
    // For most cases, a regular map with Mutex is simpler and faster
    var syncMap sync.Map

    // Store: write a key-value pair
    syncMap.Store("name", "Lucas")
    syncMap.Store("age", 25)

    // Load: read a value (returns value and whether it exists)
    if val, ok := syncMap.Load("name"); ok {
        fmt.Println("Name:", val)
    }

    // Range: iterate over all key-value pairs
    syncMap.Range(func(key, value interface{}) bool {
        fmt.Printf("  %v: %v\n", key, value)
        return true // return false to stop iteration
    })
}
