package main

import "fmt"

type person struct {
	name          string
	shengfenzheng int
	sex           string
}

type man struct {
	sex string
}

type worker struct {
	man
	person    //多继承，父类的结构体不能有相同的属性,否则使用的时候子类不知道使用那个
	deparment string
	salary    int
}

func main() {
	w := &worker{man: man{sex: "nan"},
				person: person{name: "miskak",
					           shengfenzheng: 1223213,
					           sex:           "nv"},
				deparment: "kaifa",
				salary:    2321}
	fmt.Println(w)
}
