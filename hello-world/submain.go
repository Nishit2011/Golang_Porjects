// package main

// import "fmt"

// func main() {
// 	g := greeter{
// 		greeting: "Hello",
// 		name:     "Ram",
// 	}
// 	g.greet()

// 	a := teacher{
// 		school: "dps",
// 		name:   "sita",
// 		gender: "female",
// 	}

// 	a.admit()

// }

// func (a teacher) admit() {
// 	fmt.Println(a.name, a.school)
// }

// type teacher struct {
// 	school string
// 	name   string
// 	gender string
// }

// type greeter struct {
// 	greeting string
// 	name     string
// }

// func (g greeter) greet() {
// 	fmt.Println(g.greeting, g.name)
// }
