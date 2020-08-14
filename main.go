package main

import (
	"fmt"
)
func p(i int) int {
	fmt.Println("in p",i)
	i=i+10
	return i
}
func a() {
	i := 0
	defer fmt.Println("in a",p(i))
	i++
	return
}

func c() (i int) {
	defer func() { i++ }()
	return 1
}

func main()  {
	defer func() {
		if v:=recover();v!=nil{
			fmt.Println(v)
		}
		fmt.Println("recover")
	}()
	aa()
	fmt.Println("end")
	select {

	}
}

func aa()  {

	bb()
}

func bb()  {
	 cc()
}

func cc()  {
	panic("painc")
}