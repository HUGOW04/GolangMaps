package main

import "fmt"

func main() {
	myPets := map[string]int{
		"dog":    1,
		"cat":    0,
		"cactus": 2,
	}
	fmt.Println(myPets["dog"])
	myPets["dog"] = myPets["dog"] - 1
	fmt.Println(myPets["dog"])
}
