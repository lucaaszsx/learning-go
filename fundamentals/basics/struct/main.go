package main

import "fmt"

// Person is a custom data type that groups related fields together
type Person struct {
    Name string
    Age  int
    Email string
}

// method with value receiver: does not modify the original struct
func (p Person) Greet() string {
    return fmt.Sprintf("Hi, I'm %s and I'm %d years old", p.Name, p.Age)
}

// method with pointer receiver: can modify the original struct
func (p *Person) SetAge(age int) {
    p.Age = age
}

func main() {
    // Creating a struct using field names (recommended for clarity)
    person := Person{
        Name:  "Lucas",
        Age:   25,
        Email: "lucas@example.com",
    }
    fmt.Println("Person:", person)
    fmt.Println("Name:", person.Name)

    // Creating a struct using positional values (order matters)
    person2 := Person{"John", 30, "john@example.com"}
    fmt.Println("Person2:", person2)

    // Using a method with value receiver
    fmt.Println(person.Greet())

    // Using a method with pointer receiver to modify the struct
    person.SetAge(26)
    fmt.Println("Age after SetAge:", person.Age) // 26

    // Pointer to a struct
    p := &Person{Name: "Alice", Age: 28, Email: "alice@example.com"}
    // Go allows accessing fields directly through a pointer (no need for (*p).Name)
    fmt.Println("Pointer struct name:", p.Name)

    // Struct embedding (composition over inheritance)
    employee := Employee{
        Person: Person{Name: "Bob", Age: 35, Email: "bob@example.com"},
        Company: "Tech Corp",
        Salary:  75000,
    }
    fmt.Println("Employee:", employee)
    // Embedded fields are accessed directly
    fmt.Println("Employee name:", employee.Name) // same as employee.Person.Name
    fmt.Println(employee.Greet()) // inherited method from Person
}

// Employee embeds Person, inheriting its fields and methods
type Employee struct {
    Person  // embedded struct (no field name)
    Company string
    Salary  float64
}
