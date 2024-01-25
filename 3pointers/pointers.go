package main

import "fmt"

func main() {
	age := 32          // regular var
	agePointer := &age // not named thus normally  type = *int
	// fmt.Println("age address", agePointer)
	fmt.Println("age val using dereferencing", *agePointer)
	adultYears := getAdultYears(agePointer)
	fmt.Println(adultYears)
}

func getAdultYears(age *int) int { //receiving a pointer
	return *age - 18 //dereferencing to allow maths
}
