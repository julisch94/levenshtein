# Levenshtein Distance

[![Build Status](https://travis-ci.org/julisch94/levenshtein.svg?branch=master)](https://travis-ci.org/julisch94/levenshtein)
[![Coverage Status](https://coveralls.io/repos/github/julisch94/levenshtein/badge.svg?branch=master)](https://coveralls.io/github/julisch94/levenshtein?branch=master)
[![GoDoc](https://godoc.org/github.com/julisch94/levenshtein?status.svg)](https://godoc.org/github.com/julisch94/levenshtein)

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
