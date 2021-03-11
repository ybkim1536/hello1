package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/ybkim1536/greeting1"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	message := greeting1.Hello("Gladys")
	fmt.Println(message, rand.Intn(10))
	message1 := greeting1.Goodby("Gladys")
	fmt.Println(message1, rand.Intn(10))

}
