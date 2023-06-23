/*
package main

import "fmt"

func main() {
    fmt.Println("Hello World!")
}
*/

// variables
/*
package main

import "fmt"

func main() {
    var x float64 = 20.0
    fmt.Println(x)
    fmt.Printf("x is of type %T\n", x)
    // Late Initialization
    var num int
    var amount float32
    var isValid bool
    var name string
    num = 20
    amount = 49.99
    isValid = true
    name = "Bappy"
    fmt.Println(num, amount, isValid, name)
}
*/

/* type infarence: declare variables by value:
package main

  import "fmt"

  func main() {
      // Declare and Initialize Variables without DataType
      var num = 20
      var amount = 49.99
      var isValid = true
      var name = "Bappy"
      fmt.Println(num, amount, isValid, name)
     //Using the shorthand syntax
      y := 42
      fmt.Println(y)
      fmt.Printf("y is of type %T\n", y) //y is of type int
      // Declare Multiple Variables
      num1, num2 := 20, 30          //int variables
      amount, name = 49.99, "Bappy" // float and string variables
      fmt.Println(num1, num2, amount, name)
}
*/

/* mixed variable declaration

package main

import "fmt"

func main() {
    var a, b, c = 3, 4, "foo"
    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)
    fmt.Printf("a is of type %T\n", a)
    fmt.Printf("b is of type %T\n", b)
    fmt.Printf("c is of type %T\n", c)
}

*/

/*  declare variables without value
package main

import "fmt"

func main() {
    var num int
    var amount float32
    var isValid bool
    var name string
    fmt.Println(num, amount, isValid, name) // output: 0 0 false
}
/* The string value is not clear in the output because this
is showing an empty string ("") without the double quotation */

/*
package main

import "fmt"

func main() {
    var_1, var_2 := 1, "hi"
    //declaring for the first time
    fmt.Println(var_1, var_2)
    var_3, var_2 := 2, "hello"
    //same scope, the variable is reassigned
    fmt.Println(var_3, var_2)
    if var_4, var_2 := 3, "hey"; var_4 > var_1 {
        // variable is declared again in the scope of the if condition
        fmt.Println(var_4, var_2)
    }
    fmt.Println(var_2)
    //fmt.Println(var_4) //uncommenting this should give an compilation error
}
*/
/*
package main

import "fmt"

func main() {
	a := 2
	fmt.Println(a)

	if c := 4; c>a {
		a=5
	}
	fmt.Println(a)
}
*/

/* global variable
package main

import "fmt"
// Declaring Global variable
var scalar string = "Welcome to scalar"

func main() {
    // Printing Global variable in main
    fmt.Println(scalar + " from main")
    //function call
    displayGreeting()
}

func displayGreeting() {
    // Printing Global variable from a random function
    fmt.Println(scalar + " from function")
}
*/

/* constants can never be reinitialized in golang
package main

import "fmt"

func main() {
    const myConst = 10
    fmt.Println(myConst)
    // cannot assign a new value to a constant
    // myConst = 20 // this will cause a compile-time error
}
*/

/* format specifiers
package main
import "fmt"

func main() {
   str := "Hello, world!"
   num1 := 42
   num2 := 3.14159
   num3 := 1234567890
   boolean := true
   char := 'A'
    // %T: print type of the value
    fmt.Printf("%T\n", str)
    // %v: print default format for each type
    fmt.Printf("%v\n", str)
   // String format specifiers
   fmt.Printf("%s\n", str)         // %s: print string
   fmt.Printf("%q\n", str)         // %q: print quoted string
   fmt.Printf("%x\n", []byte(str)) // %x: print hex encoding of bytes
   fmt.Printf("%X\n", []byte(str)) // %X: print uppercase hex encoding of bytes
   // Integer format specifiers
   fmt.Printf("%d\n", num1) // %d: print decimal integer
   fmt.Printf("%b\n", num1) // %b: print binary integer
   fmt.Printf("%o\n", num1) // %o: print octal integer
   fmt.Printf("%x\n", num1) // %x: print hex encoding of integer
   fmt.Printf("%X\n", num1) // %X: print uppercase hex encoding of integer
   fmt.Printf("%c\n", char) // %c: print character
   // Floating-point format specifiers
   fmt.Printf("%f\n", num2) // %f: print floating-point number
   fmt.Printf("%e\n", num2) // %e: print scientific notation of floating-point       number
   fmt.Printf("%E\n", num2) // %E: print scientific notation of floating-point number with uppercase E
   fmt.Printf("%g\n", num2) // %g: print floating-point number in decimal or scientific notation, depending on the value
   fmt.Printf("%G\n", num2) // %G: print floating-point number in decimal or scientific notation, depending on the value with uppercase E
   // Width and precision
   fmt.Printf("|%5d|\n", num1)    // %5d: print decimal integer with minimum width of 5 characters
   fmt.Printf("|%-5d|\n", num1)   // %-5d: print decimal integer with minimum width of 5 characters, left-justified
   fmt.Printf("|%5.2f|\n", num2)  // %5.2f: print floating-point number with minimum width of 5 characters and 2 digits after the decimal point
   fmt.Printf("|%-5.2f|\n", num2) // %-5.2f: print floating-point number with minimum width of 5 characters and 2 digits after the decimal point, left-justified
   // Boolean format specifiers
   fmt.Printf("%t\n", boolean) // %t: print boolean value
   // Pointer format specifier
   fmt.Printf("%p\n", &num3) // %p: print pointer address
}
*/

/* numeric type
package main

import (
    "fmt"
)

func main() {
    // Signed integers
    var a int8 = 127
    var b int16 = 32767
    var c int32 = 2147483647
    var d int64 = 9223372036854775807
    // Unsigned integers
    var e uint8 = 255
    var f uint16 = 65535
    var g uint32 = 4294967295
    var h uint64 = 18446744073709551615
    // Print the values
    fmt.Println("Signed Integers:")
    fmt.Printf("int8: %d\n", a)
    fmt.Printf("int16: %d\n", b)
    fmt.Printf("int32: %d\n", c)
    fmt.Printf("int64: %d\n", d)
    fmt.Println("Unsigned Integers:")
    fmt.Printf("uint8: %d\n", e)
    fmt.Printf("uint16: %d\n", f)
    fmt.Printf("uint32: %d\n", g)
    fmt.Printf("uint64: %d\n", h)
}
*/

/* floats 
package main

import "fmt"

func main() {
    // float32 example
    var myFloat32 float32 = 3.14
    fmt.Printf("myFloat32 = %f, type = %T\n", myFloat32, myFloat32)
    // float64 example
    var myFloat64 float64 = 3.141592653589793238462643383279502
    fmt.Printf("myFloat64 = %f, type = %T\n", myFloat64, myFloat64)

}
*/

/* complex number
package main

import "fmt"

func main() {
    c1 := complex(2, 3)
    c2 := 4 + 5i             // complex initializer syntax a + ib
    c3 := c1 + c2            // addition just like other variables
    fmt.Println("Add: ", c3) // prints "Add: (6+8i)"
    re := real(c3)           // get real part
    im := imag(c3)           // get imaginary part
    fmt.Println(re, im)      // prints 6 8
}
*/

/* byte
package main

import (
    "fmt"
)

func main() {
    // assign by decimal integer (base 10)
    var ch byte = 90
    fmt.Printf("%08b %02x %v\n", ch, ch, ch)
    // assign an ASCII character
    var ch0 byte = 'Z'
    fmt.Printf("%08b %02x %v\n", ch0, ch0, ch0)
    // assign by different bases:
    var ch1 byte = 0b01011010 // Binary
    var ch2 byte = 0o132      // Octal
    var ch3 byte = 0x5a       // heX
    fmt.Printf("%08b %02x %v\n", ch2, ch2, ch2)
    fmt.Printf("%08b %02x %v\n", ch1, ch1, ch1)
    fmt.Printf("%08b %02x %v\n", ch3, ch3, ch3)
}

// Result
// 01011010 5a 90
// 01011010 5a 90
// 01011010 5a 90
// 01011010 5a 90
// 01011010 5a 90
*/

/* rune
// Simple Go program to illustrate
// how to create a rune
package main

import (
    "fmt"
    "reflect"
)

func main() {

    // Creating a rune
    rune1 := 'B'
    rune2 := 'g'
    rune3 := '\a'
    // Displaying rune and its type
    fmt.Printf("Rune 1: %c; Unicode: %U; Type: %s", rune1, rune1, reflect.TypeOf(rune1))
    fmt.Printf("\nRune 2: %c; Unicode: %U; Type: %s", rune2, rune2, reflect.TypeOf(rune2))
    fmt.Printf("\nRune 3: Unicode: %U; Type: %s", rune3, reflect.TypeOf(rune3))

}
*/

package main

import (
    "fmt"
    "strings"
)

func main() {
    // string example
    var myString string = "Hello, World!"
    fmt.Println(myString)
    // concatenation
    str1 := "Hello, "
    str2 := "World!"
    result := str1 + str2
    fmt.Println(result) // Output: Hello, World!
    // slicing
    myString = "Hello, World!"
    fmt.Println(myString[0:5]) // Output: Hello
    // indexing
    fmt.Println(myString[7]) // Output: W
    // length of a string
    fmt.Println(len(myString)) // Output: 13
    // converting a string to uppercase
    fmt.Println(strings.ToUpper(myString)) // Output: HELLO, WORLD!
    // converting a string to lowercase
    fmt.Println(strings.ToLower(myString)) // Output: hello, world!
    // splitting a string
    str := "one,two,three,four,five"
    splitStr := strings.Split(str, ",")
    fmt.Println(splitStr) // Output: [one two three four five]
}