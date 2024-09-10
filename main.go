package main

import "fmt"
import "runtime"

func main() {
	runtime.GOMAXPROCS(2)
	var messages = make(chan string)
	var sayHelloTo = func(who string) {
		var data = fmt.Sprintf("hello %s", who)
		fmt.Println("selesai kirim data ", who)
		messages <- data
		fmt.Println("test")
	}

	go sayHelloTo("john wick")
	go sayHelloTo("ethan hunt")
	go sayHelloTo("jason bourne")

	var message1 = <-messages
	fmt.Println("message 1 ", message1)

	fmt.Println("executed")
	var message2 = <-messages
	fmt.Println("message 2", message2)

	var message3 = <-messages
	fmt.Println("message 3", message3)

}
