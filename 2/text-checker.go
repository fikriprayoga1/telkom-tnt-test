package main

import (
	"fmt"
	"log"
)

func main() {
	var text string
	var text2 string
	var mResult bool

	log.Println("Please type first input")
	fmt.Scanln(&text)

	log.Println("Please type second input")
	fmt.Scanln(&text2)

	mResult = isChangeOnce(text, text2)
	log.Println(mResult)

}

func isChangeOnce(input, input2 string) bool {
	var totalText int
	var totalText2 int
	var differentiation int
	var s []rune
	var s2 []rune

	totalText = len(input)
	totalText2 = len(input2)
	differentiation = totalText - totalText2

	if (differentiation < -1) || (differentiation > 1) {
		return false
	} else {
		s = []rune(input)
		s2 = []rune(input2)

		if totalText < totalText2 {

			for i := 0; i < totalText; i++ {
				for x := 0; x < len(s2); x++ {

					if s[i] == s2[x] {

						s2 = delChar(s2, x)
						break
					}
				}

			}

			if len(s2) > 1 {
				return false
			} else {
				return true
			}

		} else {

			for i := 0; i < totalText2; i++ {
				for x := 0; x < len(s); x++ {
					if s2[i] == s[x] {
						s = delChar(s, x)
						break
					}
				}

			}

			if len(s) > 1 {
				return false
			} else {
				return true
			}
		}

	}
}

func delChar(s []rune, index int) []rune {
	return append(s[0:index], s[index+1:]...)
}
