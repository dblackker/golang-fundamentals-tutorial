package main

import (
	"fmt"
	"math/rand"
	"time"
	"strings"
)

func main() {

	for timer := 10; timer >= 0; timer-- {
		if timer == 0 {
			fmt.Println("Boom!")
			break
		}

		fmt.Println(timer)
		time.Sleep(1 * time.Second)
	}

	//_, err := os.Open("~\\workspace\\hello.txt")
	//if err != nil {
	//	fmt.Println("Error returned was nil", err)
	//}

	//switch tmpNum := random(); tmpNum {
	//case 0,2,4,6,8,10:
	//	fmt.Println("Even number, yay!")
	//case 1,3,5,7,9:
	//	fmt.Println("Odd number, boo!")
	//default:
	//	fmt.Println("What number? That's weird.")
	//}

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

func random() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(10)
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
