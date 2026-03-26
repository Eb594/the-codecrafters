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


