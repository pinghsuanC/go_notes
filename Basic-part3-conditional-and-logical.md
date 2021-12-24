# **CONDITIONAL STATEMENTS**

- [**CONDITIONAL STATEMENTS**](#conditional-statements)
  - [**IF, IF ELSE, ELSE IF**](#if-if-else-else-if)
  - [**SWITCH STATEMENTS** - basically the same, but breaks are implied, and allows multiple cases on one branch](#switch-statements---basically-the-same-but-breaks-are-implied-and-allows-multiple-cases-on-one-branch)
    - [switch cases with **MULTIPLE TESTS** - each seperate case needs to be unique](#switch-cases-with-multiple-tests---each-seperate-case-needs-to-be-unique)
    - [**TAGLESS SYNTAX** - allowing comparison in switch](#tagless-syntax---allowing-comparison-in-switch)
    - [**FALLINGTHROUGH** - Bypass implied break](#fallingthrough---bypass-implied-break)
    - [**TYPE SWITCH**](#type-switch)
- [**COMPARISON VARIABLES** - normal](#comparison-variables---normal)
- [**LOGICAL OPERATORS** - normal](#logical-operators---normal)
- [Loop : **FOR statements**](#loop--for-statements)
  - [LABEL LOOP](#label-loop)
  - [FOR RANGE LOOP](#for-range-loop)

## **IF, IF ELSE, ELSE IF**

- You are not alloed to have single-line executions, you have to have {}
- Operators

  - ```
      if false {// execution if flase}
    ```
  - You can **directly use the result of functions**

    - Here, ok was passed and used as if condition ![if1](imgs/if1.PNG)

- if-else
- else-if

## **SWITCH STATEMENTS** - basically the same, but breaks are implied, and allows multiple cases on one branch

- simple cases
  - ![switch1](imgs/switch1.PNG)
- the target case is called a `tag`
- There is no 'break'. It's **implied**. So thoughtful. If you want to leave early, then you can still use this keyword

### switch cases with **MULTIPLE TESTS** - each seperate case needs to be unique

- Golang allows multiple cases in one switch branch
  - ![switch2](imgs/switch2.PNG)
- Note that **the test cases must be unique**
- It is allowed to use functions on switch target

  - ![switch3](imgs/switch3.PNG)

### **TAGLESS SYNTAX** - allowing comparison in switch

- When the tag case isn't there, you can switch using the comparison operators.etc
- It's allowed to overlap in range. The **first** case that's evaluated to true will be considered as success
- ![switch4](imgs/switch4.PNG)

### **FALLINGTHROUGH** - Bypass implied break

- If you want to bypass the implied break, use `fallthrough` ![switch5](imgs/switch5.PNG)
- It will fall to the next case and execute
- Note that `fallthrough` is **logicless**

### **TYPE SWITCH**

- Grab the type from the interface or whatever variable, and then we can use golang type to switch cases ![switch6](imgs/switch6.PNG)

# **COMPARISON VARIABLES** - normal

- These does't work with String type or reference type
- `<, > ==, <=, >=, !=`
- floating variables and == don't get along

# **LOGICAL OPERATORS** - normal

- `&&, ||, !(not)`
- short circuiting: the sequence of evaluation matters. If it satisbfiys already then there is no need to look forward

# Loop : **FOR statements**

- ![for5](imgs/for5.PNG)
- has **continue** and **break** keywords
- allow i++.etc
- simple loop

  ```
   for i:=0, i < 5, i++ {
     fmt.Println(i)
   }

  ```

- Golang allows to initialize multiple counters

  ```

  for i,j:=0, i < 5, i,j=i+1, j+1 {
    fmt.Println(i)
  }

  ```

- To i in the main function:

  ```
  i:=0
  for , i < 5, i, i++{
    fmt.Println(i)
  }

  ```

- missing the inrementer may cause infinite loop
- working collection with for loop

## LABEL LOOP

- By this syntax, you can label a loop and then brak by calling the label ![for1](imgs/for2.PNG)

## FOR RANGE LOOP

- this works with **collections**, string, and **channels**
- loop through and get the individual values of a slice ![for1](imgs/for1.PNG)
  - note: since go reuqires you to use all vairables, if you don't need k use \_ to replace it
- work with string ![for3](imgs/for3.PNG)
  - notice that when printing out string it gives a unicode interpretation. To solve that type convert to string ![for4](imgs/for4.PNG)
- Channel
  - used for concurrent programming
