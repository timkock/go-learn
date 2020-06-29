package main

import "fmt"

func zerobyval(ival int) {
  ival = 0
}

func zerobyptr(iptr *int) {
  *iptr = 0
}

func main() {
  i := 1
  fmt.Println("initial:", i)

  zerobyval(i)
  // pases by value hence doesn't change original variable
  fmt.Println("zeroval:", i)

  zerobyptr(&i)
  // passes by pointer (reference) and dereferences the pointer
  fmt.Println("zeroptr:", i)

  // &i syntax gives memory address of i (pointer to i)
  fmt.Println("pointer:", &i)
}
