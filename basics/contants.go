package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "세계"
	fmt.Println("Hello", World)
	fmt.Println("Happ", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
