package main

import "fmt"

func demo(i int) (func(),func()) {
	a:= func() {
		defer func(i int) {
			fmt.Println("a",i)
		}(i)
		i++
	}
	b:= func() {
		fmt.Println("b",i)
	}
	return a,b
}

func main()  {
	a,b:=demo(10)
	a()
	b()
}

