package main

import (
    "fmt"

    "reflect"
)

func main() {

    var a [5]int
    fmt.Println("emp:", a)

    a[4] = 100
    fmt.Println("set:", a)
    fmt.Println("get:", a[4])

    fmt.Println("len:", len(a))

    // defined array literal
    b := [5]int{1, 2, 3, 4, 5}
    fmt.Println("dcl(b) value:", b)
    fmt.Println("dcl(b) type kind:", reflect.TypeOf(b).Kind())

    // compiler counted array literal
    y := [...]int{1, 2, 3, 4, 5}
    fmt.Println("dcl(y) value:", y)
    fmt.Println("dcl(y) type kind:", reflect.TypeOf(y).Kind())

    // slice
    z := []int{1, 2, 3, 4, 5}
    fmt.Println("dcl(z) value:", z)
    fmt.Println("dcl(z) type kind:", reflect.TypeOf(z).Kind())

    var twoD [2][3]int
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}
