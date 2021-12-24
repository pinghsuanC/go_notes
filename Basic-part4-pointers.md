# **POINTERS** - like those in C/C++

- [**POINTERS** - like those in C/C++](#pointers---like-those-in-cc)
  - [Creation](#creation)
  - [**DEREFERENCE** pointers using **\***](#dereference-pointers-using-)
  - [pointer arithmatic is **NOT ALLOWED** - unlike C](#pointer-arithmatic-is-not-allowed---unlike-c)
  - [the **`new`** function - use it with struct](#the-new-function---use-it-with-struct)
  - [working with **`nil`**](#working-with-nil)
  - [Types with internal pointers](#types-with-internal-pointers)

## Creation

- long way of creating pointer ![pointer1](./imgs/pointer1.png)
  - the numeric representation of the address

## **DEREFERENCE** pointers using **\***

- ![pointer2](./imgs/pointer2.png)
- you can use the dereference to change the value sotred at that address
  - ![pointer3](./imgs/pointer3.png)
  - ![pointer4](./imgs/pointer4.png)

## pointer arithmatic is **NOT ALLOWED** - unlike C

- ![pointer5](./imgs/pointer5.png)
- If you want it, then use `unsafe` package

## the **`new`** function - use it with struct

- Use it with struct
  - ![pointer6](./imgs/pointer6.png)
  - The `&` means that it stroes the address of a struct that has one field with value 42 in it
- with the pointer, you can initialize it with `new()` function. All fields will be initialized to the zero-value of the type
  - ![pointer7](./imgs/pointer7.png)
- To get the value of the structs

## working with **`nil`**

- The zero-value of pointer is`nil`, it will be pinted out as `<nil>`

## Types with internal pointers
