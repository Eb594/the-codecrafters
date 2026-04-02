# variables
* variables are used for storing the data values (container)

## go variable types
* In Go, there are different types of variables, for example:

* int- stores integers (whole numbers), such as 123 or -123
* float32- stores floating point numbers, with decimals, such as 19.99 or -19.99
* string - stores text, such as "Hello World". String values are surrounded by double quotes
* bool- stores values with two states: true or false

## In Go, there are two ways to declare a variable:
1. With the var keyword:this means using the var key followed by the variable name and type
## example
* var variablename type = value
* 2. With the := sign:
Use the := sign, followed by the variable value

# Variable Declaration With Initial Value

* If the value of a variable is known from the start, you can declare the variable and assign a value to it on one line

# Example

package main

import ("fmt")

func main() {
  var student1 string = "John" 
  var student2 = "Jane" 
  x := 2 

  fmt.Println(student1)
  fmt.Println(student2)
  fmt.Println(x)
}

# Variable Declaration Without Initial Value

* In Go, all variables are initialized. So, if you declare a variable without an initial value, its value will be set to the default value of its type:
Example
package main
import ("fmt")

func main() {
  var a string
  var b int
  var c bool

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
}

# Value Assignment After Declaration

It is possible to assign a value to a variable after it is declared. This is helpful for cases the value is not initially known.
# Example
package main

import ("fmt")

func main() {
  var student1 string
  student1 = "John"
  fmt.Println(student1)
}

## Go Multiple Variable Declaration

In Go, it is possible to declare multiple variables on the same line.
## Example

This example shows how to declare multiple variables on the same line:
package main
import ("fmt")

func main() {
  var a, b, c, d int = 1, 3, 5, 7

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
  fmt.Println(d)
}

* Note: If you use the type keyword, it is only possible to declare one type of variable per line.

If the type keyword is not specified, you can declare different types of variables on the same line:
## Example
package main
import ("fmt")

func main() {
  var a, b = 6, "Hello"
  c, d := 7, "World!"

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
  fmt.Println(d)
}

## Go Variable Declaration in a Block

Multiple variable declarations can also be grouped together into a block for greater readability:
## Example
package main
import ("fmt")

func main() {
   var (
     a int
     b int = 1
     c string = "hello"
   )

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
}

## Go Variable Naming Rules

A variable can have a short name (like x and y) or a more descriptive name (age, price, carname, etc.).

## Go variable naming rules:

    A variable name must start with a letter or an underscore character (_)
    A variable name cannot start with a digit
    A variable name can only contain alpha-numeric characters and underscores (a-z, A-Z, 0-9, and _ )
    Variable names are case-sensitive (age, Age and AGE are three different variables)
    There is no limit on the length of the variable name
    A variable name cannot contain spaces
    The variable name cannot be any Go keywords

## Multi-Word Variable Names

* Variable names with more than one word can be difficult to read.

* There are several techniques you can use to make them more readable:
Camel Case

* Each word, except the first, starts with a capital letter:
* myVariableName = "John"
Pascal Case

* Each word starts with a capital letter:
* MyVariableName = "John"
Snake Case

* Each word is separated by an underscore character:
my_variable_name = "John"
