package main

import (
	"fmt"
	"math/rand"
)

func ifElseExample(){
	// IF-ELSE
	x := rand.Intn(100)
	oddEvenCheck(x)
	x = rand.Intn(5)
	oddEvenCheck(x)
	
	// IF-ELSE-IF
	
	if num := rand.Intn(50); num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }
	
	
}

func oddEvenCheck(x int ){
	if x%2 == 0 {
        fmt.Println(x," is even")
    } else {
        fmt.Println(x," is odd")
    }
}

