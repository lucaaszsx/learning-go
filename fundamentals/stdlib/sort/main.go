package main

import (
    "fmt"
    "sort"
)

func main() {
    // ----- Sorting basic types

    // Ints: sort a slice of integers
    numbers := []int{5, 2, 8, 1, 9, 3}
    fmt.Println("Before sort:", numbers)
    sort.Ints(numbers)
    fmt.Println("After sort:", numbers) // [1 2 3 5 8 9]

    // Strings: sort a slice of strings alphabetically
    names := []string{"Charlie", "Alice", "Bob", "David"}
    fmt.Println("Before sort:", names)
    sort.Strings(names)
    fmt.Println("After sort:", names) // [Alice Bob Charlie David]

    // Float64s: sort a slice of floats
    prices := []float64{9.99, 5.49, 12.99, 3.99}
    sort.Float64s(prices)
    fmt.Println("Sorted prices:", prices)

    // ----- Check if a slice is already sorted
    fmt.Println("Is sorted?", sort.IntsAreSorted(numbers)) // true

    // ----- Searching in a sorted slice
    // SearchInts returns the index where the value would be inserted
    // The slice MUST be sorted before searching
    idx := sort.SearchInts(numbers, 5)
    fmt.Println("Index of 5:", idx) // 3

    idx = sort.SearchInts(numbers, 4) // 4 is not in the slice
    fmt.Println("Index where 4 would go:", idx) // 3 (between 3 and 5)

    // ----- Reverse sorting
    // sort.Reverse wraps a sort.Interface to reverse the order
    sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
    fmt.Println("Reverse sorted:", numbers) // [9 8 5 3 2 1]

    // ----- Custom sorting with sort.Slice
    // Useful when you need to sort by a specific field or custom logic
    people := []struct {
        Name string
        Age  int
    }{
        {"Alice", 30},
        {"Bob", 25},
        {"Charlie", 35},
        {"David", 25},
    }

    // Sort by age, then by name for same age
    sort.Slice(people, func(i, j int) bool {
        if people[i].Age == people[j].Age {
            return people[i].Name < people[j].Name // secondary sort by name
        }
        return people[i].Age < people[j].Age // primary sort by age
    })

    fmt.Println("People sorted by age:")
    for _, p := range people {
        fmt.Printf("  %s (age %d)\n", p.Name, p.Age)
    }

    // ----- sort.SliceStable: preserves the original order of equal elements
    scores := []struct {
        Player string
        Score  int
    }{
        {"Alice", 100},
        {"Bob", 85},
        {"Charlie", 100},
        {"David", 90},
    }

    sort.SliceStable(scores, func(i, j int) bool {
        return scores[i].Score > scores[j].Score // descending by score
    })

    fmt.Println("Scores (stable sort, descending):")
    for _, s := range scores {
        fmt.Printf("  %s: %d\n", s.Player, s.Score)
    }
}
