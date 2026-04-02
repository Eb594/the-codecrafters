## Go Functions

* A function is a block of statements that can be used repeatedly in a program.

* A function will not execute automatically when a page loads.

* A function will be executed by a call to the function.
Create a Function

* To create (often referred to as declare) a function, do the following:

    * Use the func keyword.
    Specify a name for the function, followed by parentheses ().
    Finally, add code that defines what the function should do, inside curly braces {}.

## Call a Function

* Functions are not executed immediately. They are "saved for later use", and will be executed when they are called.

* In the example below, we create a function named "myMessage()". The opening curly brace ( { ) indicates the beginning of the function code, and the closing curly brace ( } ) indicates the end of the function. The function outputs "I just got executed!". To call the function, just write its name followed by two parentheses ():
## Example
package main
import ("fmt")

func myMessage() {
  fmt.Println("I just got executed!")
}

func main() {
  myMessage() // call the function
}

## Result:
I just got executed!

A function can be called multiple times.
Example
package main
import ("fmt")

func myMessage() {
  fmt.Println("I just got executed!")
}

func main() {
  myMessage()
  myMessage()
  myMessage()
}

Result:
I just got executed!
I just got executed!
I just got executed!


## Naming Rules for Go Functions

* A function name must start with a letter
* A function name can only contain alpha-numeric characters and underscores (A-z, 0-9, and _ )
* Function names are case-sensitive
* A function name cannot contain spaces
* If the function name consists of multiple words, techniques introduced for multi-word variable naming can be used

## Go Function Parameters and Arguments
* Parameters and Arguments

* Information can be passed to functions as a parameter. Parameters act as variables inside the function.

* Parameters and their types are specified after the function name, inside the parentheses. You can add as many parameters as you want, just separate them with a comma:
Syntax
func FunctionName(param1 type, param2 type, param3 type) {
  // code to be executed
}
## Function With Parameter 
## Example

* The following example has a function with one parameter (fname) of type string. When the familyName() function is called, we also pass along a name (e.g. Liam), and the name is used inside the function, which outputs several different first names, but an equal last name:
## Example
package main
import ("fmt")

func familyName(fname string) {
  fmt.Println("Hello", fname, "Refsnes")
}

func main() {
  familyName("Liam")
  familyName("Jenny")
  familyName("Anja")
}

## Result:
Hello Liam Refsnes
Hello Jenny Refsnes
Hello Anja Refsnes

* When a parameter is passed to the function, it is called an argument. So, from the example above: fname is a parameter, while Liam, Jenny and Anja are arguments.

## Multiple Parameters

* Inside the function, you can add as many parameters as you want:
## Example
package main
import ("fmt")

func familyName(fname string, age int) {
  fmt.Println("Hello", age, "year old", fname, "Refsnes")
}

func main() {
  familyName("Liam", 3)
  familyName("Jenny", 14)
  familyName("Anja", 30)
}

## Result:
Hello 3 year old Liam Refsnes
Hello 14 year old Jenny Refsnes
Hello 30 year old Anja Refsnes

* When you are working with multiple parameters, the function call must have the same number of arguments as there are parameters, and the arguments must be passed in the same order.

## Return Values

* If you want the function to return a value, you need to define the data type of the return value (such as int, string, etc), and also use the return keyword inside the function:
Syntax
func FunctionName(param1 type, param2 type) type {
  // code to be executed
  return output
}
## Function Return Example
## Example

Here, myFunction() receives two integers (x and y) and returns their addition (x + y) as integer (int):
package main
import ("fmt")

func myFunction(x int, y int) int {
  return x + y
}

func main() {
  fmt.Println(myFunction(1, 2))
}

## Result:
3
## Named Return Values

* In Go, you can name the return values of a function.
## Example

* Here, we name the return value as result (of type int), and return the value with a naked return (means that we use the return statement without specifying the variable name):
package main
import ("fmt")

func myFunction(x int, y int) (result int) {
  result = x + y
  return
}

func main() {
  fmt.Println(myFunction(1, 2))
}

## Result:
3

* The example above can also be written like this. Here, the return statement specifies the variable name:
## Example
package main
import ("fmt")

func myFunction(x int, y int) (result int) {
  result = x + y
  return result
}

func main() {
  fmt.Println(myFunction(1, 2))
}

* Store the Return Value in a Variable

* You can also store the return value in a variable, like this:
## Example

* Here, we store the return value in a variable called total:
package main
import ("fmt")

func myFunction(x int, y int) (result int) {
  result = x + y
  return
}

func main() {
  total := myFunction(1, 2)
  fmt.Println(total)
}

## Multiple Return Values

* Go functions can also return multiple values.
## Example

* Here, myFunction() returns one integer (result) and one string (txt1):
package main
import ("fmt")

func myFunction(x int, y string) (result int, txt1 string) {
  result = x + x
  txt1 = y + " World!"
  return
}

func main() {
  fmt.Println(myFunction(5, "Hello"))
}

## Result:
10 Hello World!
## Example

* Here, we store the two return values into two variables (a and b):
package main
import ("fmt")

func myFunction(x int, y string) (result int, txt1 string) {
  result = x + x
  txt1 = y + " World!"
  return
}

func main() {
  a, b := myFunction(5, "Hello")
  fmt.Println(a, b)
}

## Result:
10 Hello World!

* If we (for some reason) do not want to use some of the returned values, we can add an underscore (_), to omit this value.
## Example

* Here, we want to omit the first returned value (result - which is stored in variable a):
package main
import ("fmt")

func myFunction(x int, y string) (result int, txt1 string) {
  result = x + x
  txt1 = y + " World!"
  return
}

func main() {
   _, b := myFunction(5, "Hello")
  fmt.Println(b)
}

## Result:
Hello World!
