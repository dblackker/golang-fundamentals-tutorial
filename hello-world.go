package main

import (
	"fmt"
	"strings"
)

func main() {

	name := "Daniel"
	course := "Docker Deep Dive"

	fmt.Println(converter(name, course))

	//fmt.Println("\nHi", name, "you're currently watching",
	//	course)

	//changeCourse(&course)

	//fmt.Println("\nYou are now watching course", course)
}

func changeCourse(course *string) string {

	*course = "First Look: React Native"
	fmt.Println("\nTrying to change your course to", *course)

	return *course
}

func converter(s1, s2 string) (str1, str2 string) {
	s1 = strings.Title(s1)
	s2 = strings.ToUpper(s2)

	return s1, s2
}
