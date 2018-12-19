package main

import (
	"fmt"
)

// var rmdirs []func()
// for _, d := range tempDirs() {
// 	dir := d
// 	os.MkdirAll(dir, 0755)
// 	rmdirs = append(rmdirs, func() {
// 		os.RemoveAll(dir)
// 	})
// }

// for _, rmdir := range rmdirs {
// 	rmdir()
// }

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}
func f(...int) {}
func g([]int)  {}

func main7() {

	fmt.Printf("%T\n", f)
	fmt.Printf("%T\n", g)

	// fmt.Println(sum())
	// fmt.Println(sum(3))
	// fmt.Println(sum(1, 2, 3, 4))
}
