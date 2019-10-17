package main

import (
	"fmt"

	"github.com/Abaygithub/GoLangBasic/002_package/nested/hello"
	"github.com/Abaygithub/GoLangBasic/002_package/nested/say"
)

func main() {

	fmt.Println("Start program...")

	fmt.Println(say.CallFromSay() + hello.CallFromHello())

}
