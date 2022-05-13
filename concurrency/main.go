// package main

// import "fmt"

// func main() {
// 	myChannel := make(chan int)
// 	go sendValues(myChannel)

// 	// for i := 10; i < 25; i++ {
// 	// 	value, open := <-myChannel
// 	// 	if !open {
// 	// 		break
// 	// 	}
// 	// 	fmt.Println(value)
// 	// 	// fmt.Println(<-myChannel)
// 	// }

// 	for value := range myChannel {
// 		fmt.Println(value)
// 	}

// }

// func sendValues(myChannel chan int) {

// 	for i := 10; i < 20; i++ {
// 		myChannel <- i
// 	}
// 	close(myChannel)
// }
