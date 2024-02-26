package main

import (
	"fmt"
	"math/rand"
	"time"
)

var h int = 0
var b int = 0
var bt int = 0

func main() {
	TrinkenTimer()
	HungerTimer()
}
func LebenTimer() {
	h--
	if h <= 2 {
		fmt.Println("ich sterbe")
	}
	if h == 0 {
		fmt.Println("ich bin tot")
	}
	fmt.Println(h)
}
func TrinkenTimer() {
	for i := 0; i < bt; bt-- {
		if bt <= 2 {
			fmt.Println("ich habe durst")
		}
		if bt == 0 {
			LebenTimer()
		}
		Rand := rand.Intn(180) + 180
		time.Sleep(time.Duration(Rand) * time.Second)
		fmt.Println(bt)
	}

}
func HungerTimer() {
	for i := 0; i < b; b-- {
		if b <= 2 {
			fmt.Println("ich habe hunger")
		}
		if bt == 0 {
			LebenTimer()
		}
		Rand := rand.Intn(150) + 30
		time.Sleep(time.Duration(Rand) * time.Second)
		fmt.Println(b)
	}

}
