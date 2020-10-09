package main

import (
	"fmt"
	"moshi/moshi_instance"
)

func main() {
	s := singleton.Instance.Add()
	fmt.Println(s)
	s = singleton.Instance.Add()
	fmt.Println(s)
	s = singleton.Instance.Add()
	fmt.Println(s)

	se := singleton.New()
	fmt.Println(se.Add())
	se2 := singleton.New() //再次new返回同一个instance
	fmt.Println(se2.Add())
	fmt.Println(se.Add())
}
