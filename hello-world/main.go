package main

import "fmt"

type School struct {
	name       string `required max: "100"`
	department string
}

type Teacher struct {
	school   School
	number   int
	subject  string
	students []string
}

func main() {
	// var str string = "Nishit is a good boy"

	// var principal float64 = 3000
	// var rate float64 = 2.5
	// var time float64 = 2
	// var simpleInterest = (principal * (rate) * time) / 100
	// fmt.Printf("%v %T\n", simpleInterest, simpleInterest)
	// mapExample := make(map[string]string)
	// mapExample = map[string]string{
	// 	"Bihar":     "Patna",
	// 	"Jharkhand": "Ranchi",
	// }
	// mapExample["Karnataka"] = "Bengaluru"
	// mapExample["Maharashtra"] = "Mumbai"
	// delete(mapExample, "Jharkhand")
	// pop, ok := mapExample["Bihar"]
	// fmt.Println(pop, ok)

	// aTeacher := Teacher{
	// 	number:   1,
	// 	subject:  "Maths",
	// 	students: []string{"Ram", "Shyam", "Sohan", "Mohan"},
	// 	school:   School{name: "KV", department: "Science"},
	// }

	//fmt.Println(aTeacher)

	//animal := struct{ name string }{name: "Ram"}
	//fmt.Println(animal)

	//fmt.Println(Teacher)
	//fmt.Printf(aTeacher.school.name)

	// bTeacher := Teacher{
	// 	number:   2,
	// 	subject:  "Geography",
	// 	students: []string{"Alex", "Mike", "Seeta", "Rizwan"},
	// 	school:   School{name: "DAV", department: "Arts"},
	// }
	// fmt.Printf(bTeacher.school.name)

	// cTeacher := Teacher{}
	// cTeacher.school.name = "DPS"
	// dTeacher := &cTeacher
	// dTeacher.school.name = ""

	// fmt.Println(dTeacher)
	// fmt.Println(cTeacher)

	//count := 9

	// if count < 10 {
	// 	fmt.Printf("less than 10")
	// }
	// if count > 100 {
	// 	fmt.Printf("greater than 100")
	// }

	// switch {
	// case count < 10:
	// 	fmt.Printf("less than 10")
	// 	fallthrough
	// case count > 100:
	// 	fmt.Printf("greater than 100")
	// }
	//arr := [...]string{}
	// var i interface{} = "1"
	// switch i.(type) {
	// case int:
	// 	fmt.Printf("Datatype is integer")
	// case string:

	// 	fmt.Printf("Datatype is string")

	// case bool:
	// 	fmt.Printf("Datatype is boolean")
	// case float64:
	// 	fmt.Printf("Datatype is float64")
	// case []string:
	// default:
	// 	fmt.Printf("Datatype is others")
	// }

	// for i, j := 0, 0; i < 5; i, j = i+1, j+i {
	// 	fmt.Printf("%v \n", i)
	// 	fmt.Println(j)
	// }

	// for i := 0; i < 10; i++ {
	// 	for j := 0; j < 10; j = j + i {
	// 		if j == 5 {
	// 			fmt.Println(j)
	// 			break
	// 		}

	// 	}

	// }

	// mapEx1 := map[string]string{
	// 	"India":     "Delhi",
	// 	"Pakistan":  "Islamabad",
	// 	"Australia": "Canberra",
	// }

	// for _, v := range mapEx1 {
	// 	fmt.Println(v)
	// }

	// fmt.Println("start")
	// panicker()
	// fmt.Println("end")

	// var a int = 42
	// var b *int = &a
	// fmt.Println(&a, *b)
	// a = 27
	// fmt.Println(&a, *b)

	// c := [3]int{1, 2, 3}
	// d := &c[1]
	// e := &c[2]

	// fmt.Printf("%p %p\n", d, e)
	// fmt.Printf("%v %v\n", *d, *e)

	// ms := struct{ name string }{name: "Rakesh"}
	// ns := &ms

	// fmt.Printf("%p \n", ns)
	// ns.name = "Mahesh"
	// fmt.Printf("%v \n", ms)

	// xs := [3]int{10, 11, 12}
	// ys := &xs
	// fmt.Printf("%v \n", *ys)

	// s := sum(1, 2)
	// fmt.Println(s)
	// greeting := "Hello"
	// name := "Nishit!"

	// getMessage(&greeting, &name)
	// fmt.Println(name)

	// s := sum(1, 2, 3, 4, 5)
	// fmt.Println(*s)

	s, err := divideTwoNumbers(1.0, 0.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s)
}

func divideTwoNumbers(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot be divided by zero")
	}
	result := a / b
	return result, nil

}

// func sum(values ...int) *int {
// 	fmt.Println(values)
// 	result := 0

// 	for _, v := range values {
// 		result = result + v
// 	}
// 	fmt.Println(&result)
// 	return &result
// }

// func getMessage(greeting, name *string) {
// 	fmt.Println(*greeting, *name)
// 	*name = "Ted"
// 	fmt.Println(*name)

// }

// func sum(a, b int) (result int) {

// 	result = a + b
// 	return result
// }

// func panicker() {
// 	fmt.Println("about to panic")
// 	defer func() {
// 		if err := recover(); err != nil {
// 			log.Println("Error:", err)

// 		}
// 	}()
// 	panic("something bads happened")
// 	fmt.Println("done panicking")

// }
