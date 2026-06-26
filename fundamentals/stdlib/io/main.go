package main

import (
    "bytes"
    "fmt"
    "io"
    "strings"
)

func main() {
    // ----- The io.Reader and io.Writer interfaces
    // These are the two fundamental interfaces in Go's I/O system
    // Reader: anything that can read bytes (files, network, strings, etc.)
    // Writer: anything that can write bytes (files, network, buffers, etc.)

    // strings.NewReader creates a Reader from a string
    reader := strings.NewReader("Hello, io package!")

    // Read: reads bytes into a buffer
    buf := make([]byte, 5)
    n, err := reader.Read(buf)
    if err != nil && err != io.EOF {
        fmt.Println("Read error:", err)
        return
    }
    fmt.Printf("Read %d bytes: %s\n", n, buf) // "Hello"

    // ReadAll: reads everything until EOF
    reader.Reset("This is the full content")
    all, err := io.ReadAll(reader)
    if err != nil {
        fmt.Println("ReadAll error:", err)
        return
    }
    fmt.Println("ReadAll:", string(all))

    // ----- bytes.Buffer: a versatile in-memory buffer
    // It implements both Reader and Writer interfaces
    var buf2 bytes.Buffer

    // Write: writes bytes to the buffer
    buf2.WriteString("Hello, ")
    buf2.WriteString("Buffer!")
    fmt.Println("Buffer content:", buf2.String())

    // Read from the buffer
    data := make([]byte, 5)
    n, _ = buf2.Read(data)
    fmt.Printf("Read from buffer: %s (%d bytes)\n", data, n)

    // ----- Copy: transfer data from Reader to Writer
    source := strings.NewReader("Data to be copied")
    var destination bytes.Buffer

    // io.Copy reads from source and writes to destination
    bytesCopied, err := io.Copy(&destination, source)
    if err != nil {
        fmt.Println("Copy error:", err)
        return
    }
    fmt.Printf("Copied %d bytes: %s\n", bytesCopied, destination.String())

    // ----- LimitReader: read only N bytes
    longReader := strings.NewReader("This is a very long string that we want to truncate")
    limited := io.LimitReader(longReader, 10) // read at most 10 bytes

    limitedData, _ := io.ReadAll(limited)
    fmt.Println("Limited read:", string(limitedData)) // "This is a "

    // ----- MultiReader: combine multiple readers into one
    r1 := strings.NewReader("First, ")
    r2 := strings.NewReader("Second, ")
    r3 := strings.NewReader("Third!")

    multi := io.MultiReader(r1, r2, r3)
    combined, _ := io.ReadAll(multi)
    fmt.Println("MultiReader:", string(combined)) // "First, Second, Third!"

    // ----- TeeReader: read and write simultaneously
    source2 := strings.NewReader("TeeReader example")
    var captured bytes.Buffer

    // TeeReader writes everything it reads into the writer (captured)
    tee := io.TeeReader(source2, &captured)
    result, _ := io.ReadAll(tee)
    fmt.Println("TeeReader result:", string(result))
    fmt.Println("Captured copy:", captured.String()) // same content!
}
