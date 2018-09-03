package main

import (
	"fmt"
	//"reflect"
	"time"
)

func switchExample() {

	whatAmI := func(i interface{}) {
		//tt := reflect.TypeOf(i)
		//fmt.Println("Got : ", tt)
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		case string:
			fmt.Println("I'm an string")
		case time.Time:
			fmt.Println("I'm an time.Time")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI(1.23)
	whatAmI("hey")
	whatAmI(time.Now())

}
