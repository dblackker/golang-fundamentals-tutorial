package main

import (
	"fmt"
	"strings"
)

func main() {

	bestFinish := bestLeagueFinishes(13, 10, 13, 17, 14, 5, 4)

	fmt.Println(bestFinish)

	//name := "Daniel"
	//course := "Docker Deep Dive"

	//fmt.Println(converter(name, course))

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

func bestLeagueFinishes(finishes ...int) int {

	best := finishes[0]
	for _, i := range finishes {
		if i < best {
			best = i
		}
	}

	return best
}
