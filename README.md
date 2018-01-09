# Levenshtein Distance

Simple golang calculator for Leveshtein Distance. Using the techniques of **Clean Code**.

## Look inside

Go ahead and take a look at the code. Let me know what you think about the coding style. 

## Usage Example

You're only one import away from using my little library.

    package main

    import (
     "github.com/julisch94/levenshtein"
     "fmt"
    )

    func main() {
      fmt.Println("Should be 2: ", levenshtein.CalculateDistance("mandy", "mary"))
    }
