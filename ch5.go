package main

import (
	"fmt"
	"sort"
)

func add(x int, y int) int   { return x + y }
func sub(x, y int) (z int)   { z = x - y; return }
func first(x int, _ int) int { return x }
func zero(int, int) int      { return 0 }

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)
	visitAll =
		func(items []string) {
			for _, item := range items {
				if !seen[item] {
					seen[item] = true
					visitAll(m[item])
					order = append(order, item)
				}
			}
		}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}

func main6() {

	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d\t%s\n", i+1, course)
	}

	f := squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	// fmt.Printf("%T\n", add)
	// fmt.Printf("%T\n", sub)
	// fmt.Printf("%T\n", first)
	// fmt.Printf("%T\n", zero)
}
