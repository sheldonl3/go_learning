package main

import "fmt"

//func reverse(a, b *int) {
//	*a, *b = *b, *a
//	return
//}

func delete_map(m map[string]string, key string) {
	delete(m, key)
}

func add_map(m map[string]string,key string,value string){
	m[key]=value
}

func delete_slice(lis *[]string, index int) { //函数传递slice,有修改要传指针
	*lis = append((*lis)[:index], (*lis)[index+1:]...)
}

func main() {
	sil:=[]int{1,2,3}
	fmt.Println(sil)
	fmt.Println(sil[1:])
}
