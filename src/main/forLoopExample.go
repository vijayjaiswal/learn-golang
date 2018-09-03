package main

import (
	"fmt"	
)

func forExample(){
	// Conditional For
	i:=1
	for i <= 5 {
		fmt.Println("i :" , i)
		i++
	}
	// Simple For
	
	for y:=10 ; y<=15 ; y++ {
		fmt.Println("y :" , y)
	}
	
	// For break
	
	for z:=10 ; z<=15 ; z++ {
		fmt.Println("z :" , z)
		if z==12 {
			fmt.Println("Breaking...")
			break
		}
	}
	
	// For continue
	
	for x:=10 ; x<=15 ; x++ {
		if x%2 == 0 {
			fmt.Println("Continuing...")
			continue
		}
		fmt.Println("x :" , x)
	}
	
	
}


