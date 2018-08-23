package main

import (
	"fmt"
)

func main() {
	fmt.Printf("hello, world\n")
	 values()
	 printCurrentTime()
}

func values() {
	fmt.Println("go" + "lang")
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
