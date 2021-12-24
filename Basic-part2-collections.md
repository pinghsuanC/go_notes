# **COLLECTIONS**

- [**COLLECTIONS**](#collections)
  - [**ARRAYS** - same as an common array in Java, except that it's copying a literal copy; Can use & and \* in C](#arrays---same-as-an-common-array-in-java-except-that-its-copying-a-literal-copy-can-use--and--in-c)
    - [Creation](#creation)
    - [Built-in functions](#built-in-functions)
  - [**SLICES** - reference type, similar to List interface in Java (underlying may be an ArrayList?); Python-like slice operations](#slices---reference-type-similar-to-list-interface-in-java-underlying-may-be-an-arraylist-python-like-slice-operations)
    - [Creation](#creation-1)
    - [Built-in functions (looks like Python)](#built-in-functions-looks-like-python)
  - [**MAPS** - Similar to Maps in Java](#maps---similar-to-maps-in-java)
  - [**STRUCTS** Similar to C++ struct or class in java](#structs-similar-to-c-struct-or-class-in-java)
    - [**EMBEDDING** - Golang way of 'inheritance' by composition](#embedding---golang-way-of-inheritance-by-composition)
    - [**TAGS** - somehow reminds me of styled components](#tags---somehow-reminds-me-of-styled-components)

## **ARRAYS** - same as an common array in Java, except that it's copying a literal copy; Can use & and \* in C

- the elements in array needs to be of the same type.
- arrays are considered values. When you are copying an array, you are creating a literal copy
  - `a := [...]int{1,2,3}`
  - `b := a` here b holds a separate copy of a
- Golang do has pointers. To avoid to copy over everything, use pointers
  - `a := [...]int{1,2,3}`
  - `b := &a` here b holds an address of a

### Creation

- `grades := [3]int{97, 85, 93}`
  - length, type, initialization
- `grades := [...]int{97, 85, 93}`
  - if you know the litetal data, then you don't need to input the number. use `...` to indicate to use the number of elements input afterwards
- `var student [3]string` empty array
  - `students[0] = "test"`
- multidimention array: `[sizeDim1][sizeDim2]..[sizeLastDim]type`

### Built-in functions

- `len(Array)`

## **SLICES** - reference type, similar to List interface in Java (underlying may be an ArrayList?); Python-like slice operations

- a projection of underlying array
- Slice is a **reference type**
  - `a := [...]int{1,2,3}`
  - `b := a` here b holds a reference to a
- the capacity doesn't need to be fixed

### Creation

- `a := []int{1,2,3}` the length isn't needed at compile time
- `make` function
  - `a := make([]int, 3)`
    - The length and capacity are both set to 3
  - `a := make([]int, 3, 100)`
    - you can pass in the third argument for capacity

### Built-in functions (looks like Python)

- `cap(slice)` similar to len(array)
  - the Capacity may be different from the length of underlying array
- Functions to create 'slices' ![slice1](./imgs/slice1.png)
  - Note that the slices are still reference-type, therefore if you change the value of the original slice, the slices that contain the corresponding value will also be changed ![slice2](./imgs/slice2.png)
  - Use the these oprations to pop/remove elements
    - `b := a[:len(a)-1]` remove the last element
    - `b := a[1:]` remove the first element
    - `b := append(a[:2], a[3:]...)` remove the 3rd element; however if you print a afterwards there will be a duplicate of the last element. Make sure to remove the references if needed.
- `make(arrayType, length, capacity=len)` slice creation
- `append(slice, element1, element2, ...)` return a slice that with the input elements. Everything after the slice is going to be interpreted as a thing to append
  - If it exceeds the original capacity, it will copy over and add more capacity
  - that's why it's beeter to initialize it with make() if we know the approximate length
  - to concatenate elements, use `...` operator
    - `append(slice1, ...slice2 )`

## **MAPS** - Similar to Maps in Java

- key - value pairs
- **Important** the order of the map is not guaranteed. It's very likely that it's not maintained.
- Map is a reference type. The operation on one map will be affecting all maps
- Creation

  - Direct declaration:

    - `name := map[KEY_TYPE]VALUE_TYPE{KEY1:VALUE1...}`
    - e.g. ![map1](./imgs/map1.png)

  - all the key must be of the same type, and all the value must be of the same type
  - **Constraint: the keys MUST be able to compared for equality**. e.g. primitive types, arrays
    - The slice, therefore can **not** be a key of a map, but an array can
    - `m := map[[3]int]string{}`
  -

- Operations
  - `make(map[string]int)` use make to make empty maps
  - `make(map[string]int, 10)` the 10 isn't commonly used
  - `MAP[KEY]` returns the value associated and whether the pair was found
    - If you don't have the key inside the map, it will return value 0 of the value type (e.g. 0 if value is int; "" if value is string), which may cause some troubles. It's common pracitce to do the following:
    - ![map2](./imgs/map2.png)
    - The second return value of the `MAP[KEY]` will be false if it's not found.
  - `delete(MAP, KEY)` delete the key-value pair in the map
- `len(MAP)` return the length

## **STRUCTS** Similar to C++ struct or class in java

- ![map3](./imgs/map3.png)
  - unlike arrays and map, the fields can be any type we want
- structs are value types. When passing structs aronud, you are passing copies around. Like array, if we want to use references, use addresses and pointers.
- Usually used as type
- Creation
  - To create a struct
  ```
   type STRUCT strct{
     value_1 TYPE_1
     value_2 TYPE_2
     ...
     }
  ```
  - To create a struct instance:
  ```
    varName := STRUCT{ values }
  ```
  - To get the values, use dot syntax `varName.FieldName`
  - Positional syntax: you don't need to write out the struct field names when you are declaring an instance
    - ![struct1](./imgs/struct1.png)
    - This is usually not recommended because the order are pron to be changed.
    - The fields are generally not strictly recommended
  - Creating on-the-fly (**anonymous**)
    - ![struct2](./imgs/struct2.png)
    - first {} defines the structure of the struct, second {} defines the values of the struct
- Operations
- Naming conventions: Same as variables.
  - lowercase first letter = package level
  - uppercase first letter = export outside of the package

### **EMBEDDING** - Golang way of 'inheritance' by composition

- Technically, there is no **inheritance** in Golang, but it uses a model to realize the relationship
- By doing this, the animal struct is embeded in the bird ![struct3](./imgs/struct3.png)
- In this way, we can initialize the fields in the Animal struct in Bird ![struct4](./imgs/struct4.png)
- It "looks like" the bird is having the fields of Animla, but it's just syntaic sugar.
- You can declare the Animal struct directly if you already now the value ![struct5](./imgs/struct5.png)

### **TAGS** - somehow reminds me of styled components

- describe some specific information of a field in struct
- arglength argument: ![struct6](./imgs/struct6.png)
- If you need to use key-value pairs, you will need to use column to seperate the key and the value, and the value are typically put in quptation marks.
- They are not directly accessible. To get the requirements when you are using them, use `reflection` package
  - ![struct7](./imgs/struct7.png)
    - get type of struct => get field => get `Tag` property of the field
- There are validation libraries or frameworks that reads the tags. JSON package some times use this.
