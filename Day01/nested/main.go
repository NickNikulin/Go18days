package main

import (
	"Go18days/Day01/nested/hello"
	"Go18days/Day01/nested/say"
	"fmt"
)

func main() {
	fmt.Println("Start program...")
	fmt.Println(say.CallFromSay() + hello.CallHello())
}
