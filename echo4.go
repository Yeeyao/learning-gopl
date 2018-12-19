package main

import (
	"fmt"
)

// var s string
// var i, j, k int
// var b, f, s2 = true, 2.3, "four"

// var boiling float64 = 100
// var n = flag.Bool("n", false, "omit trailing newline")
// var sep = flag.String("s", " ", "separator")

func main3() {
	medals := []string{"gold", "silver", "bronze"}
	fmt.Println(gcd(3000, 22))
	fmt.Println(fib(10))
	// x, y := 1, 2
	// fmt.Println(x, y)
	// x, y = y, x
	// fmt.Println(x, y)
	// v := 1
	// v++
	// fmt.Println(v)
	// p := new(int)
	// q := new(int)
	// fmt.Println(p == q)
	// p := new(int)
	// fmt.Println(*p)
	// *p = 2
	// fmt.Println(*p)
	// i, j := 0, 1
	// t := 0.0
	// x := 1
	// p := &x
	// fmt.Println(*p)
	// *p = 2
	// fmt.Println(x)
	// fmt.Println(i, j)
	// fmt.Println(t)

	// var x, y int
	// fmt.Println(&x == &x, &x == &y, &x == nil)
	// var p = f()
	// fmt.Println(*p)
	// fmt.Println(f() == f())

	// flag.Parse()
	// fmt.Print(strings.Join(flag.Args(), *sep))
	// if !*n {
	// 	fmt.Println()
	// }
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

func newInt() *int {
	return new(int)
}

func newInt2() *int {
	var dummy int
	return &dummy
}

// func f() *int {
// 	v := 1
// 	return &v
// }
