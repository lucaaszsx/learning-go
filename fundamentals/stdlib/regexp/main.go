package main

import (
    "fmt"
    "regexp"
)

func main() {
    // ----- Basic pattern matching

    // MatchString: checks if a string matches a pattern
    matched, _ := regexp.MatchString(`\d+`, "Hello 123")
    fmt.Println("Contains digits:", matched) // true

    // Compile: create a reusable regexp (more efficient for repeated use)
    re, err := regexp.Compile(`[a-z]+@[a-z]+\.[a-z]+`)
    if err != nil {
        fmt.Println("Compile error:", err)
        return
    }

    // Match: check if the pattern matches
    fmt.Println("Is email?", re.MatchString("user@example.com")) // true
    fmt.Println("Is email?", re.MatchString("not-an-email"))      // false

    // ----- Finding matches

    text := "Contact us at support@example.com or sales@company.org"

    // FindString: returns the first match
    email := re.FindString(text)
    fmt.Println("First email:", email) // support@example.com

    // FindAllString: returns all matches
    // second argument is the max number of matches (-1 = all)
    emails := re.FindAllString(text, -1)
    fmt.Println("All emails:", emails) // [support@example.com sales@company.org]

    // FindStringIndex: returns the start and end positions of the first match
    idx := re.FindStringIndex(text)
    fmt.Println("First email position:", idx) // [14 33]

    // ----- Capture groups: extract parts of a match
    // Use parentheses () to define groups
    re2 := regexp.MustCompile(`(\w+)@(\w+)\.(\w+)`)
    // MustCompile is like Compile but panics on error (use when pattern is constant)

    // FindStringSubmatch: returns the full match and each group
    match := re2.FindStringSubmatch("user@example.com")
    fmt.Println("Full match:", match[0])  // user@example.com
    fmt.Println("Username:", match[1])     // user
    fmt.Println("Domain:", match[2])       // example
    fmt.Println("TLD:", match[3])          // com

    // FindStringSubmatch with names
    re3 := regexp.MustCompile(`(?P<user>\w+)@(?P<domain>\w+)\.(?P<tld>\w+)`)
    match = re3.FindStringSubmatch("admin@server.com")
    for i, name := range re3.SubexpNames() {
        if i > 0 && name != "" {
            fmt.Printf("%s: %s\n", name, match[i])
        }
    }

    // ----- Replacing

    // ReplaceAllString: replace all occurrences
    censored := re.ReplaceAllString(text, "[REDACTED]")
    fmt.Println("Censored:", censored) // "Contact us at [REDACTED] or [REDACTED]"

    // ReplaceAllStringFunc: replace using a function
    upper := re.ReplaceAllStringFunc(text, func(match string) string {
        return fmt.Sprintf("<%s>", match)
    })
    fmt.Println("Tagged:", upper)

    // ----- Split
    re4 := regexp.MustCompile(`\s*,\s*`)
    parts := re4.Split("apple, banana, cherry", -1)
    fmt.Println("Split:", parts) // [apple banana cherry]

    // ----- Common patterns
    // Phone number: (XX) XXXXX-XXXX
    phoneRe := regexp.MustCompile(`\(\d{2}\)\s*\d{5}-\d{4}`)
    fmt.Println("Valid phone?", phoneRe.MatchString("(11) 99999-1234"))

    // URL
    urlRe := regexp.MustCompile(`https?://[\w.-]+(/\S*)?`)
    urls := urlRe.FindAllString("Visit https://go.dev or http://example.com/path?q=1", -1)
    fmt.Println("URLs found:", urls)
}
