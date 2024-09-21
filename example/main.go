package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/dqx0/ulid"
)

func main() {
	simple()
	custom()
}

func simple() {
	fmt.Println(ulid.Make())
}
func custom() {
	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	ms := ulid.Timestamp(time.Now())

	ulid, err := ulid.New(ms, entropy)
	if err != nil {
		panic(err)
	}
	fmt.Println(ulid)
}
