package main

func main5() {

	type Movie struct {
		Title strng 
		Year int 'json:"released"'
		Color bool 'json:"color, omitempty"'
		Actors []string
	}

	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "ingride Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}},
		{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}}
	}

	type tree struct {
		value       int
		left, right *tree
	}

	// type Employee struct {
	// 	ID        int
	// 	Name      string
	// 	Address   string
	// 	DoB       time.Time
	// 	Position  string
	// 	Salary    int
	// 	ManagerID int
	// }

	// var dilbert Employee

	// ages := make(map[string]int)

	// ages := map[string]int{
	// 	"alice":   31,
	// 	"charlie": 34,
	// }

	// var names []string
	// for name := range ages {
	// 	names = append(names, name)
	// }

	// sort.Strings(names)

	// for _, name := range names {
	// 	fmt.Printf("%s\t%d\n", name, ages[name])
	// }

	// ages["alice"] = 32
	// fmt.Println(ages["alice"])

	// age, ok := ages["bob"]
	// if !ok {
	// 	fmt.Println(age, ok)
	// }

	// type Currency int

	// const (
	// 	USD Currency = iota
	// 	EUR
	// 	GBP
	// 	RMB
	// )

	// symbol := [...]string{USD: "$", EUR: "€", GBP: "£", RMB: "¥"}

	// fmt.Println(RMB, symbol[RMB])

	// c1 := sha256.Sum256([]byte("x"))
	// c2 := sha256.Sum256([]byte("X"))
	// fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)

	// months := [...]string{1: "January", /* ... */, 12: "December"}
	// fmt.Println(months)
	// var a [3]int
	// var q [3]int = [3]int{1, 2, 3}
	// p := [...]int{1, 2, 3}
	// var r [3]int = [3]int{1, 2}
	// fmt.Printf("%T\n", p)
	// fmt.Println(q[2])
	// fmt.Println(r[2])
	// fmt.Println(a[0])
	// fmt.Println(a[len(a)-1])

	// for i, v := range a {
	// 	fmt.Printf("%d %d\n", i, v)
	// }

	// for _, v := range a {
	// 	fmt.Printf("%d\n", v)
	// }

}

func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nill {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
