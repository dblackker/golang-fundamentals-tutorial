package main

import (
	"fmt"
	"strings"
)

func main() {

	if firstRank, secondRank := 39, 614; firstRank > secondRank {
		fmt.Println("\nFirst course is doing better")
	} else if secondRank < firstRank {
		fmt.Println("\nOh dear... your first course must be doing abysmally.")
	} else {
		fmt.Println("They're the same")
	}

	//bestFinish := bestLeagueFinishes(13, 10, 13, 17, 14, 5, 4)

	//fmt.Println(bestFinish)

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
