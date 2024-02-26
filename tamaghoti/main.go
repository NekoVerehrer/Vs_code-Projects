package main

import (
	"fmt"
	"math/rand"

	//"strconv"
	"time"
)

var h int = 10

const (
	time_b1  = 10
	time_b2  = 10
	time_bt1 = 10
	time_bt2 = 10
)

func main() {
	go TrinkenTimer(5)
	HungerTimer(5)
}
func LebenTimer() {
	h--
	if h <= 2 {
		fmt.Println("ich sterbe")
	}
	if h == 0 {
		fmt.Println("ich bin tot")
	}
	fmt.Println("Leben:", h)
}
func TrinkenTimer(bt int) {
	for bt > 0 {
		if bt <= 2 {
			fmt.Println("ich habe durst")
		}
		if bt == 0 {
			LebenTimer()
		}
		Rand := rand.Intn(time_bt2) + time_bt1
		time.Sleep(time.Duration(Rand) * time.Second)
		fmt.Println("Durst:", bt)
		bt--
	}

}
func HungerTimer(b int) {
	for b > 0 {
		if b <= 2 {
			fmt.Println("ich habe hunger")
		}
		if b == 0 {
			LebenTimer()
		}
		Rand := rand.Intn(time_b2) + time_b1
		time.Sleep(time.Duration(Rand) * time.Second)
		fmt.Println("Hunger:", b)
		b--
	}

}
