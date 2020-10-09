package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	No   int
	Name string
}

func main() {
	s := Student{1, "name"}

	m1 := map[int]string{1: "misaka", 2: "sheldon"}
	m2 := map[int]string{2: "sheldon", 1: "misaa"}
	//t := reflect.TypeOf(m1)
	//v := reflect.ValueOf(m1)
	stu_type := reflect.TypeOf(s)
	stu_value := reflect.ValueOf(s)
	fmt.Println("value:", stu_value, "type:", stu_type)
	for i := 0; i < stu_type.NumField(); i++ {
		key := stu_type.Field(i)
		value := stu_value.Field(i).Interface()
		fmt.Println( key, value)
	}
	fmt.Println(reflect.DeepEqual(m1, m2))

}
