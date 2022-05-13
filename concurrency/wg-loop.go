// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func printValue(num int, wg *sync.WaitGroup) {
// 	for i := 0; i < 5; i++ {

// 		fmt.Println(num)
// 		time.Sleep(50 * time.Millisecond)

// 	}
// 	wg.Done()
// }

// func main() {
// 	// var wg sync.WaitGroup

// 	// for i := 0; i < 5; i++ {
// 	// 	wg.Add(1)
// 	// 	//channel1 <- i
// 	// 	go printValue(i, &wg)

// 	// }
// 	// wg.Wait()
// 	channel1 := make(chan int)
// 	channel1 <- 23

// 	go func() {
// 		<-channel1
// 	}()

// 	fmt.Println(<-channel1)
// }
