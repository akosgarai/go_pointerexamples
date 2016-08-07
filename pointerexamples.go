package main

import (
    "fmt"
)

type Vertex struct {
    X int
    Y int
}

func main () {
    i, j := 42, 2701
    pointerExample(i, j)
    vertexExample(i, j)
    arrayExample()
    arrayExample_2()
}

func vertexExample (i, j int) {
    v := Vertex{i, j}
    fmt.Println(v)
    p := &v
    fmt.Println(*p)
    p.X = 0
    fmt.Println(*p)
}

func arrayExample_2 () {
    names := [4]string{
        "John",
        "Paul",
        "George",
        "Ringo",
    }
    fmt.Println(names)

    a := names[0:2]
    b := names[1:3]
    fmt.Println(a, b)

    b[0] = "XXX"
    fmt.Println(a, b)
    fmt.Println(names)
}
func arrayExample () {
    var a [2]string
    a[0] = "Hello"
    a[1] = "World"
    fmt.Println(a[0], a[1])
    fmt.Println(a)

    primes := [6]int{2, 3, 5, 7, 11, 13}
    fmt.Println(primes)
    var s []int = primes[1:4]
    fmt.Println(s)
}

func pointerExample (i, j int) {
    p := &i     //point to i
    fmt.Println(*p) // read i through the pointer
    *p = 21         // set i through the pointer
    fmt.Println(i)  // see the new value of i

    p = &j         // point to j
    *p = *p / 37   // divide j through the pointer
    fmt.Println(j) // see the new value of j
}
