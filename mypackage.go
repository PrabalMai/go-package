//mypackage.go
package mypackage

import "fmt"

// Greet greets the person with the given name.
func Greet(name string) string {
    return fmt.Sprintf("Hello, %s!", name)
}

// Add adds two integers and returns the result.
func Add(a, b int) int {
    return a + b
}
