package day4

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

func Part1(input string) int {

	for number := 0; number < 10_000_000; number++ {
		data := input + strconv.Itoa(number)
		hashSum := md5.Sum([]byte(data))
		hashString := fmt.Sprintf("%x", hashSum)

		if hashString[0:5] == "00000" {
			return number
		}
	}
	return -1
}

func Part2(input string) int {
	for number := 0; number < 100_000_000; number++ {
		data := input + strconv.Itoa(number)
		hashSum := md5.Sum([]byte(data))
		hashString := fmt.Sprintf("%x", hashSum)

		if hashString[0:6] == "000000" {
			return number
		}
	}
	return -1
}
