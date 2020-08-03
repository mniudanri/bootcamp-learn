package main

import (
    "fmt"
)

func main () {
  fmt.Println("sum: ", 3 + 2)      // sum int typed
  fmt.Println("sum: ", 3.0 + 2.0)  // sum float64 typed
  fmt.Println("sum: ", 3 + 0.5)    // sum float64 typed
  fmt.Println("dif: ", 7 - .5)     // difference float64 typed
  fmt.Println("dif: ", 8.5 - 1)    // difference float64 typed

  fmt.Println("prod: ", 10 * .1)    // product float64 typed
  fmt.Println("quot: ", 10 / 3)     // quotient int typed (decimal floor)
  fmt.Println("quot: ", 10 / 2)     // quotient int typed

  fmt.Println("quot: ", 10 / float64(3))  // quotient float64 typed
}
