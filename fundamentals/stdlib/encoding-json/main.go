package main

import (
    "encoding/json"
    "fmt"
)

// Struct tags define how fields are encoded/decoded in JSON
// `json:"field_name"` sets the JSON key name
// `json:"field_name,omitempty"` omits the field if it's the zero value
// `json:"-"` completely ignores the field
type User struct {
    Name    string `json:"name"`
    Age     int    `json:"age"`
    Email   string `json:"email"`
    Active  bool   `json:"active"`
    Address string `json:"address,omitempty"` // omitted if empty
    Secret  string `json:"-"`                 // never serialized
}

func main() {
    // ----- Marshal: convert Go struct to JSON bytes

    user := User{
        Name:   "Lucas",
        Age:    25,
        Email:  "lucas@example.com",
        Active: true,
        Secret: "hidden",
    }

    // json.Marshal returns []byte and an error
    jsonBytes, err := json.Marshal(user)
    if err != nil {
        fmt.Println("Marshal error:", err)
        return
    }
    fmt.Println("JSON:", string(jsonBytes))
    // {"name":"Lucas","age":25,"email":"lucas@example.com","active":true}
    // Note: "address" is omitted (empty), "secret" is excluded (json:"-")

    // MarshalIndent: same but with pretty formatting
    pretty, _ := json.MarshalIndent(user, "", "  ")
    fmt.Println("Pretty JSON:")
    fmt.Println(string(pretty))

    // ----- Unmarshal: convert JSON bytes to Go struct

    jsonStr := `{"name":"Alice","age":30,"email":"alice@example.com","active":false}`
    var parsed User

    // json.Unmarshal needs a pointer to the target variable
    err = json.Unmarshal([]byte(jsonStr), &parsed)
    if err != nil {
        fmt.Println("Unmarshal error:", err)
        return
    }
    fmt.Printf("Parsed: %+v\n", parsed)
    fmt.Println("Name:", parsed.Name)

    // ----- Working with maps (when structure is unknown)

    jsonMap := `{"city":"São Paulo","country":"Brazil","population":12300000}`
    var data map[string]interface{} // generic map for any JSON object

    err = json.Unmarshal([]byte(jsonMap), &data)
    if err != nil {
        fmt.Println("Map unmarshal error:", err)
        return
    }
    fmt.Println("City:", data["city"])
    fmt.Println("Population:", data["population"])

    // ----- Encoding slices

    fruits := []string{"apple", "banana", "cherry"}
    fruitsJSON, _ := json.Marshal(fruits)
    fmt.Println("Fruits JSON:", string(fruitsJSON)) // ["apple","banana","cherry"]

    // Decoding a JSON array
    var decoded []string
    json.Unmarshal(fruitsJSON, &decoded)
    fmt.Println("Decoded fruits:", decoded)

    // ----- JSON with nested structs

    type Address struct {
        Street string `json:"street"`
        City   string `json:"city"`
    }

    type Person struct {
        Name    string  `json:"name"`
        Address Address `json:"address"`
        Tags    []string `json:"tags"`
    }

    person := Person{
        Name:    "Bob",
        Address: Address{Street: "123 Main St", City: "Portland"},
        Tags:    []string{"developer", "gopher"},
    }

    personJSON, _ := json.MarshalIndent(person, "", "  ")
    fmt.Println("Nested JSON:")
    fmt.Println(string(personJSON))

    // ----- Streaming: Encoder and Decoder
    // More efficient for large data or network I/O
    // json.NewEncoder(writer).Encode(value) - writes JSON to a writer
    // json.NewDecoder(reader).Decode(&value) - reads JSON from a reader
    // These work with any io.Writer/io.Reader (files, HTTP bodies, etc.)
}
