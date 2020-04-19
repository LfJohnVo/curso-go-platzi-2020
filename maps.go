package main

import "fmt"

func main(){
	n1 := make(map[string]int)

	n1["a"] = 8

	fmt.Println(n1)
	fmt.Println(n1["a"])
}
