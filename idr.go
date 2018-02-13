package main

import (
	"math/rand"
)







func main() {

	lawrence := &User{
		Username: "lawrence",
		Password: "password",
		Id: rand.Int(),
		Score: 0,
		Languages: []Language{TYPESCRIPT, GO},
		Bucket: []Idea{},
	}



}
