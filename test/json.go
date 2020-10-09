package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Book struct {
	id      int    `json:"id"` //结构体大写，key才能正确的序列化
	Title   string `json:"title"`
	Summary string `json:"summary"`
	Author  string `json:"author"`
}

func main() {
	var books map[int]*Book = make(map[int]*Book)
	book1 := Book{id: 1, Title: "c programmger", Summary: "none", Author: "misaka"}
	book2 := Book{id: 2, Title: "golang", Author: "sheldon"}
	books[book1.id] = &book1
	books[book2.id] = &book2
	data, err := json.Marshal(books)
	if err != nil {
		fmt.Println("error in marshal")
		return
	}
	err = ioutil.WriteFile("books.json", data, 0644)
	if err != nil {
		fmt.Println("error in svae file")
		return
	}
	fmt.Println("read from file")
	data, _ = ioutil.ReadFile("books.json")
	var data_books map[int]*Book
	json.Unmarshal(data, &data_books)
	for _, each := range data_books {
		fmt.Printf("%#v\n", each)
	}
}
