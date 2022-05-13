// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func WelcomeMessage() {
// 	fmt.Println("Welcome to Golang!!!")
// }

// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(2)
// 	go func() {
// 		WelcomeMessage()
// 		wg.Done()
// 	}()

// 	go func() {
// 		fmt.Println("thanks for msg")
// 		wg.Done()
// 	}()
// 	wg.Wait()
// }
