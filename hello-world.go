package main

import (
	"fmt"
)

func main() {

	name := "Daniel"
	course := "Docker Deep Dive"

	fmt.Println("\nHi", name, "you're currently watching",
		course)

	changeCourse(&course)

	fmt.Println("\nYou are now watching course", course)
}

func changeCourse(course *string) string {

	*course = "First Look: React Native"
	fmt.Println("\nTrying to change your course to", *course)

	return *course
}
