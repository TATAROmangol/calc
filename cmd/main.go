package main

import (
	"fmt"

	"example.com/m/pkg/calc"
) 



func main(){
	res := "2+(1))"
	err := calc.CalcQuots(&res)
	fmt.Println(res)
	fmt.Println(err)
}