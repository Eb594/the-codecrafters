# outputs
* go output refers to the result of gotten after or your desired result being printed 
# examples
* Go has three functions to output text:

* Print()
* Println()
* Printf()

# The Print() Function

* The Print() function prints its arguments with their default format.
## Example

Print the values of i and j:

package main
import ("fmt")

func main() {
  var i,j string = "Hello","World"

  fmt.Print(i)
  fmt.Print(j)
}

# The Println() Function

* The Println() function is similar to Print() with the difference that a whitespace is added between the arguments, and a newline is added at the end:
#  Example

package main
import ("fmt")

func main() {
  var i,j string = "Hello","World"

  fmt.Println(i,j)
}

# Printf()

Go offers several formatting verbs that can be used with the Printf() function.

Verb 	Description
%v 	||	Prints the value in Go-syntax format
%T |	Prints the type of the value
%% |	Prints the % sign
  

  ...

