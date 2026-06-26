package main

import (
    "fmt"
    "strings"
)

func main() {
    str := "Hello, World!"

    // Contains: checks if a substring exists
    fmt.Println("Contains 'World':", strings.Contains(str, "World")) // true
    fmt.Println("Contains 'world':", strings.Contains(str, "world")) // false (case-sensitive)

    // HasPrefix and HasSuffix: check the start/end of a string
    fmt.Println("HasPrefix 'Hello':", strings.HasPrefix(str, "Hello"))
    fmt.Println("HasSuffix '!':", strings.HasSuffix(str, "!"))

    // ToUpper and ToLower: case conversion
    fmt.Println("ToUpper:", strings.ToUpper(str))
    fmt.Println("ToLower:", strings.ToLower(str))

    // Trim: remove characters from both ends
    padded := "   Hello, World!   "
    fmt.Println("TrimSpace:", strings.TrimSpace(padded)) // removes leading/trailing spaces
    fmt.Println("Trim '!':", strings.Trim("!!Hello!!", "!")) // removes specific chars

    // Replace: substitute occurrences
    fmt.Println("Replace:", strings.Replace("foo bar foo", "foo", "baz", 1)) // replace first occurrence
    fmt.Println("ReplaceAll:", strings.ReplaceAll("foo bar foo", "foo", "baz")) // replace all

    // Split and Join: break and combine strings
    csv := "apple,banana,cherry"
    parts := strings.Split(csv, ",")
    fmt.Println("Split:", parts)

    joined := strings.Join(parts, " | ")
    fmt.Println("Join:", joined)

    // Index: find the position of a substring
    fmt.Println("Index of 'World':", strings.Index(str, "World")) // 7
    fmt.Println("Index of 'xyz':", strings.Index(str, "xyz"))     // -1 (not found)

    // Count: count occurrences
    fmt.Println("Count 'l':", strings.Count(str, "l")) // 3

    // Repeat: repeat a string N times
    fmt.Println("Repeat:", strings.Repeat("Go!", 3))

    // Fields: split by whitespace (useful for parsing words)
    sentence := "  Go   is   awesome  "
    words := strings.Fields(sentence)
    fmt.Println("Fields:", words) // ["Go", "is", "awesome"]
}
