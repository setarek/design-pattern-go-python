package main

import (
	"fmt"
	"sync"
)

var instance *Singleton
var once sync.Once

type Singleton struct{}

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}

func main() {
	s1 := GetInstance()
	s2 := GetInstance()

	fmt.Println(s1 == s2)
}
