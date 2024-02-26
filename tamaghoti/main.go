package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Tamagotchi struct {
	Name   string
	Hunger int
	Thirst int
	Life   int
}

const (
	time_b1  = 10
	time_b2  = 10
	time_bt1 = 10
	time_bt2 = 10
)

func main() {
	fmt.Println("Welcome to Tamagotchi!")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your Tamagotchi's name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	t := NewTamagotchi(name)
	go t.HungerTimer()
	go t.ThirstTimer()
	for {
		fmt.Println("\n\nWhat would you like to do?")
		fmt.Println("1. Feed")
		fmt.Println("2. Give a drink")
		fmt.Println("3. View Status")
		fmt.Println("4. Quit")

		fmt.Print("Enter your choice: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			t.Feed()
			fmt.Println("You feed your Tamagotchi.")
		case "2":
			t.Give_a_drink()
			fmt.Println("You give your Tamagotchi a drink.")
		case "3":
			t.PrintStatus()
		case "4":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

// Options: --------------------------------------------------------------------------------------------

func NewTamagotchi(name string) *Tamagotchi {
	return &Tamagotchi{
		Name:   name,
		Hunger: 5,
		Thirst: 5,
		Life:   10,
	}
}
func (t *Tamagotchi) Feed() {
	t.Hunger++
	if t.Hunger >= 0 {
		t.Hunger = 5
	}
}
func (t *Tamagotchi) Give_a_drink() {
	t.Thirst++
	if t.Thirst >= 5 {
		t.Thirst = 5
	}
}
func (t *Tamagotchi) Life_down() {
	t.Thirst--
	if t.Life <= 2 {
		fmt.Println("ich sterbe")
	}
	if t.Thirst <= 0 {
		t.Thirst = 0
		fmt.Println("ich bin tot")
	}
}
func (t *Tamagotchi) PrintStatus() {
	fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\nNext:3")
	fmt.Println("----------------------------------------------------------------------")
	fmt.Printf("Name: %s\n", t.Name)
	fmt.Printf("Hunger: %d\n", t.Hunger)
	fmt.Printf("Thirst: %d\n", t.Hunger)
	fmt.Printf("Life: %d\n", t.Hunger)
}

// TIMER: ---------------------------------------------------------------------------------------------

func (t *Tamagotchi) ThirstTimer() {
	for t.Life > 0 {
		if t.Thirst > 0 {
			t.Thirst--
			if t.Thirst <= 2 {
				fmt.Println("ich habe durst")
			}
			Rand := rand.Intn(time_bt2) + time_bt1
			time.Sleep(time.Duration(Rand) * time.Second)
		} else {
			// Random time to wait between time_bt1 - (time_bt2 + time_bt1)
			Rand := rand.Intn(time_bt2) + time_bt1
			time.Sleep(time.Duration(Rand) * time.Second)
			t.Life_down()
		}
	}
}
func (t *Tamagotchi) HungerTimer() {
	for t.Life > 0 {
		if t.Hunger > 0 {
			t.Hunger--
			if t.Hunger <= 2 {
				fmt.Println("ich habe hunger")
			}
			// fmt.Println("Hunger:", b)
			Rand := rand.Intn(time_b2) + time_b1
			// fmt.Println(Rand)
			time.Sleep(time.Duration(Rand) * time.Second)
		} else {
			// Random time to wait between time_b1 - (time_b2 + time_b1)
			Rand := rand.Intn(time_b2) + time_b1
			time.Sleep(time.Duration(Rand) * time.Second)
			t.Life_down()
		}
	}
}
