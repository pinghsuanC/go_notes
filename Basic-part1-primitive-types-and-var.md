## [Jump to main page](./main.md)

# Primitive types, variables and constants

- [Primitive types, variables and constants](#primitive-types-variables-and-constants)
- [**PRIMITIVE TYPES**](#primitive-types)
  - [**BOOLEAN** - default value is 'false', not zero](#boolean---default-value-is-false-not-zero)
  - [**NUMERIC TYPES & OPERATIONS** - default value is 0; int, float and complex](#numeric-types--operations---default-value-is-0-int-float-and-complex)
    - [**Integers** - signed(8, 16, 32, 64); unsigned(8, 16, 32)](#integers---signed8-16-32-64-unsigned8-16-32)
    - [**Floats** - IEEE-754 standard; float(32, 64)](#floats---ieee-754-standard-float32-64)
    - [**Complex Number** - using float](#complex-number---using-float)
  - [Text types](#text-types)
    - [**STRING** - immutable, UTF8 character, an allias of Byte](#string---immutable-utf8-character-an-allias-of-byte)
    - [**RUNE** - UTF32 character, using '', alias of int32, need special package](#rune---utf32-character-using--alias-of-int32-need-special-package)
- [**VARIABLES**](#variables)
  - [Variable declaration](#variable-declaration)
  - [Redeclaration and shadowing, visibility](#redeclaration-and-shadowing-visibility)
  - [Naming Conventions](#naming-conventions)
  - [**Naming Conventions controls Visibility**](#naming-conventions-controls-visibility)
  - [Type conversions](#type-conversions)
- [**CONSTANTS**](#constants)
  - [Naming convention](#naming-convention)
  - [**TYPED CONSTANTS** - similar to typed variables, just aren't mutable](#typed-constants---similar-to-typed-variables-just-arent-mutable)
  - [**UNTYPED CONSTANTS** - rely on compiler's type inference, may have implicit conversion](#untyped-constants---rely-on-compilers-type-inference-may-have-implicit-conversion)
  - [**ENUMERATED CONSTANTS (IOTA) & ENUMERATION EXPRESSIONS** - similar to Java Enumerate type, more functional though; need to be available at compile time](#enumerated-constants-iota--enumeration-expressions---similar-to-java-enumerate-type-more-functional-though-need-to-be-available-at-compile-time)

# **PRIMITIVE TYPES**

## **BOOLEAN** - default value is 'false', not zero

- `var n bool = true`
- `var n bool = false`
- `n:= 1 == 1`
- `m:= 1 == 2` common use case
- everytime you initialize a variable in go, it actually has a zero value. The zero value for go is `false`, therefore by default it's false.
  - ![primitiveType1](imgs/primitiveType1.PNG)
- Note that the boolean isn't using -1 and 0.etc like js does

## **NUMERIC TYPES & OPERATIONS** - default value is 0; int, float and complex

- the zero value is 0, or the 0 of the specific type

### **Integers** - signed(8, 16, 32, 64); unsigned(8, 16, 32)

- **signed int** of unspecified size (32 or 64 depending on system)
  - usually the default tpye
  - **int8** -128 ~ 127
  - **int16** -32768 ~ 32767
  - **int32** -2147483648 ~ 2147483647
  - **int64** -9223372036854775808 ~ 9223372036854775807
- **Unsigned int**
  - `var n uint16 = 42`
  - **unint8** 0 ~ 255
  - **int16** 0 ~ 65535
  - **int32** 0 ~ 4294967295
- built-in operations
  - `+, -, *, /, %`
  - note these are not changing the types implicitly without explicit conversion
- built-in bit-operators (similar to C++)
  - `&` AND
  - `|` OR
  - `^` exclusive OR
  - `&^` AND NOT
    - true if neighther of the bit has the bit set
  - Example ![numericOperation1](imgs/numericOperation1.PNG)
- built-in bit-shifter (similar to C++)
  - `<<`
  - `>>`
  - Example ![numericOperation2](imgs/numericOperation2.PNG)

### **Floats** - IEEE-754 standard; float(32, 64)

- IEEE-754 standard
- by default, initialzing symbol always use `float64`
- `float32` (+-)1.18E^-38 ~ (+-)3.4E^38
- `float64` (+-)2.23E^-308 ~ (+-)1.8E^308
- note that these needs conversion
- Creating floating numbers: allow the expression of e(float32) and E(float64)
  - ![floatNumber1](imgs/floatNumber1.PNG)
- Operations
  - `+, -, *, /`
  - there is no reminder (`%`) operator

### **Complex Number** - using float

- `var n complex64 = 1 + 2i` -> float32 + float32
- `var n complex128 = 1 + 2i` -> float64 + float64
- Operations
  - `+, -, *, /`
  - To decompose the parts, there are two built-in functions
    - ![complexNumber1](imgs/complexNumber1.PNG)
    - real() and imag()
  - To create Mega-complex number
    - ![complexNumber2](imgs/complexNumber2.PNG)

## Text types

### **STRING** - immutable, UTF8 character, an allias of Byte

- In Go, it stands for any UTF8 character, and is an allias of Byte
- generally immutable
  - We usually can't change the character
  - ![string3](imgs/string3.PNG)
- Note that String can't encode every type that's availabe
- It's similar to lower-level languages where it's treated like an array of characters
  - asking for the third letter directly, it will give back an integer(uint8): ![string1](imgs/string1.PNG)
  - To get the string back, we need type conversion: ![string2](imgs/string2.PNG)
- Operations
  - String concatenation
    - ![string4](imgs/string4.PNG)
  - Converting to a collection of bytes (`byte slices`), which is more flexible
    - ![string5](imgs/string5.PNG)

### **RUNE** - UTF32 character, using '', alias of int32, need special package

- a **true** type alias. When printing out the type, the type will be `int32`
- Represents an `UTF32` characters
- werid type, UTF32 can be up to 32 bits long, but not need to
- Declaration:
  - `r := 'a' ` Note that it uses single quote
  - ![rune1](imgs/rune1.PNG)
  - It's an alias of int32, rather than a byte collection like string (rune===int32)
- use special functions from strings package from Golang

---

# **VARIABLES**

## Variable declaration

1. Declaring in a function
   - normal declarations & value assignment
     - `var i int`
     - variable i with type integer
     - Assign value by `i = 42`
   - Assign within the same line / full declaration
     - `var i int = 42`
   - **Short cut for decalaration and value assign**
     - `i:= 42`
     - note that we don't need `var` here, since go is doing both declaration and type inference
     - `j:= 42.` put a dot at the end if you need it to be float64. However it's not possible to initialize a float32 with this syntax
2. Declaring at package level
   - you have to use the full declaration `var i int = 42`
   - you can create a block of variables (like struct)
     - ![variableBlock](imgs/variableBlock.PNG)
     - There can be multiple variable blocks. Better organization!

## Redeclaration and shadowing, visibility

- You are not allowed to redeclare a vairable twice in the same scope
  - ![variableRedeclaration1](imgs/variableRedeclaration1.PNG)
- However it's allowed to shadowing. The declaration and value in the innermost scope is taking precidense

  - ![variableRedeclaration2](imgs/variableRedeclaration2.PNG)

- All variables **MUST** be used. If you have a local variable that is declared but not used, that is a **compiled-time error**

## Naming Conventions

- Pascal or camelCase, and as short as reasonable
- basically, use meaningful name if the variable is going to be around for a while or it's going to be used outside of the package
- it's better to kep all acronyms in uppercase for readability
  - e.g. theUrl => `the URL`
  - e.g. theHttp => theHTTP

## **Naming Conventions controls Visibility**

- Naming controls visibility

- There are three levels of visibility in Go
  - package-level scope
    - **lower case variables** means that this variable is scoped to this package
    - everthing that's consuming the package can't see or work with it
  - global level scope
    - **upper case variables** means it's exported to this package, and it exposes this variable to the outside package
  - block scope
    - e.g. variable declared inside a function
    - never visible to the outside of the scope {}

## Type conversions

- similar to Java, usually there is no implicit type conversion. It's better practice to make type conversion explicit.
  - ![typeconversion1](imgs/typeConversion1.PNG)
- note that there may be information lost during type conversion
  - `destinationType(variable)`
- Need to understand the string with Go when converting to String
  - ![typeconversion2](imgs/typeConversion2.PNG)
  - in Go, the string is an allas for a string of bytes
  - in the screenshot, the conversoin is reads "asign value 4 to a type char" and 42 translates to "\*"
  - to convert into String, use `strconv` package
  - ![typeconversion3](imgs/typeConversion3.PNG)

# **CONSTANTS**

- not allowed to change the values
- has to assignable at _compile time_.
  - E.g. you can't use Math.Sin(1.57) to assign a constant since the result is returned after execution

## Naming convention

- `const myConst` or `const MyConst` depending on whether you want to export the variables

## **TYPED CONSTANTS** - similar to typed variables, just aren't mutable

- similar to typed variables
- you can use any primitive type in Go
- `const myConst int = 42`
- note that collectoin types (e.g array) are inherently mutable, therefore you can't create constants with them.
- They can be shadowed both in value and type
  - ![const1](imgs/const1.PNG)
- Operations are the same to primitive variables

## **UNTYPED CONSTANTS** - rely on compiler's type inference, may have implicit conversion

- use compiler's ability to infer the type for us
  - ![const2](imgs/const2.PNG)
  - it will replace all instances of a
  - implicit conversion!

## **ENUMERATED CONSTANTS (IOTA) & ENUMERATION EXPRESSIONS** - similar to Java Enumerate type, more functional though; need to be available at compile time

- Most commonly applied at the package level
- `const a = iota`
- iota is a counter when we are creating a enumerate constant
- create a cont block, as it executes it will assign values to each variable
  - ![const3](imgs/const3.PNG)
- if we stop assigning iota after the first one, the compiler is going to infer the pattern of assignment
  - ![const5](imgs/const5.PNG)
- If we create other blocks, then they will start over with 0
  - ![const6](imgs/const6.PNG)
- note that since it starts with 0, it may collide with the default value of int. To solve this:
  1. use the first iota as error value
     - ![const7](imgs/const7.PNG)
  2. use `_` variable, the Golang only write-only variable
     - ![const8](imgs/const8.PNG)
     - it means we don't care and throw it away after compilation, and don't use it
- It's actually allows primitive operation in const
  - ![const9](imgs/const9.PNG)
  - it can change the offsset of the
- Math.exp() won't be availble at compile time. But if you want to do exponentional, you can do it by bit-shifting.
  - ![const10](imgs/const10.PNG)
  - iota wil increment bytes as well

---

[Jump to main page](./main.md)
