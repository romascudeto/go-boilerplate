package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var messages = make(chan string)

//Channeling ... practice Channeling
func Channeling(c *gin.Context) {

	go sayHelloTo("john wick")
	go sayHelloTo("ethan hunt")
	go sayHelloTo("jason bourne")

	var message1 = <-messages
	fmt.Println(message1)

	var message2 = <-messages
	fmt.Println(message2)

	var message3 = <-messages
	fmt.Println(message3)
	fmt.Println("===========")
}

func sayHelloTo(who string) {
	var data = fmt.Sprintf("hello %s", who)
	messages <- data
}

//Channeling2 ... practice Channeling2
func Channeling2(c *gin.Context) {

	for _, each := range []string{"wick", "hunt", "bourne"} {
		go func(who string) {
			var data = fmt.Sprintf("hello %s", who)
			messages <- data
		}(each)
	}

	for i := 0; i < 3; i++ {
		printMessage(messages)
	}
}

func printMessage(what chan string) {
	fmt.Println(<-what)
}
