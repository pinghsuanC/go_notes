# Table of Content

- [Table of Content](#table-of-content)
- [Credit](#credit)
- [Jumpy to notes](#jumpy-to-notes)
  - [Basics](#basics)
  - [Something _special_ about Go](#something-special-about-go)
- [Running and package](#running-and-package)
- [Random notes](#random-notes)

# Credit

- Following [this awsome youtube tutorial by freeCodeCamp](https://www.youtube.com/watch?v=YS4e4q9oBaU&t=4752s) for the basics of golang

# Jumpy to notes

## Basics

- [Primitive Types and Variables](./Basic-part1-primitive-types-and-var.md)
- [Collections](./Basic-part2-collections.md)
- [Conditions and Logics](./Basic-part3-conditional-and-logical.md)
- [Pointers](./Basic-part4-pointers.md)

## Something _special_ about Go

- [Function Flow Control](./Function-Control.md)
- [Interface](./Interface.md)
- [Goroutines](./Goroutines.md)
- [Channels](./Channels.md)

# Running and package

- basic structure:

  - contains bin, pkg and src
  - ![basicProjectStructure](imgs/basicProjectStructure.PNG)

- `go run PATH/\*.go`
  - compile temporarily and run
- `go build PACKAGE_PATH`
  - if found a main function & it will give an executable
- `go install`
  - expected to point to a package tat has an entry point, and install that executable into the bin folder

---

# Random notes

- printf() substituters
  - `%v` value
  - `%t` type
  - `%p` pointer (will print out address)
- There are no function overloading or default values for arguments, and it's by design
- There is no generic types so repetition is possible
- Error handling can be annoying
- Everything in Go is passed by value: `passing an int value to a function makes a copy of the int, and passing a pointer value makes a copy of the pointer, but not the data it points to.`
- Official tutorial: [Effective Go](https://go.dev/doc/effective_go)
- This is an interesting read about the motivation of designs: [Go FAQ](https://go.dev/doc/faq)
