package main

import (
	"math/rand"

	"github.com/icrowley/fake"
)

func give_id() string {
	words := ""
	for i := 0; i < 3; i++ {
		words += fake.Word() + "-"
	}
	sequence := "1q2w3e4r5t6y7u8i9o0opasdfghjklzxcvbnm"
	id := ""
	for i := 0; i < 7; i++ {
		id += string(sequence[rand.Intn(len(sequence))])
	}

	id = words + id
	println(id)
	return id
}
