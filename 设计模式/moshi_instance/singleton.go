package singleton

import "sync"

//饿汉式单例模式
type singleton struct {
	count int
}

var Instance = new(singleton)

func (s *singleton) Add() int {
	s.count++
	return s.count
}

// 懒汉式单例模式
var instance *singleton


var once sync.Once

func New() *singleton {
	once.Do(func() {
		instance = new(singleton)
	})

	return instance
}
