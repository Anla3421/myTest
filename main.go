package main

import (
	"fmt"
	"time"
)

func writeData(intChan chan int) {
	for {
		intChan <- 123
		fmt.Printf("writeData = %v\n", 123)
		time.Sleep(time.Second)
	}
}

// func writeData2(secChan chan int) {
// 	for i := 1; i <= 3; i++ {
// 		secChan <- i
// 		fmt.Printf("writeSecData ==> %v\n", i)
// 	}
// }
func main() {
	intChan := make(chan int)
	// secChan := make(chan int, 10)
	// go writeData(intChan)
	go writeData(intChan)
	// go writeData2(secChan)
	// for {
	// 	select {
	// 	case v := <-intChan:
	// 		fmt.Printf("ReadData = %v\n", v)

	// 	default:
	// 		fmt.Printf("Nothine input\n")
	// 		time.Sleep(10 * time.Second)
	// 	}

	// }
	for {
		select {
		case v := <-intChan:
			fmt.Printf("ReadData = %v\n", v)
		}
	}
	fmt.Println("end process")

}
