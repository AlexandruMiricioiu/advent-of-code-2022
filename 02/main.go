package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func sum(arr []int) int {
	acc := 0

	for _, v := range arr {
		acc += v
	}

	return acc
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	s := strings.Split(string(content), "\n")

	t := 0
	c := make([]int, 0)
	for _, el := range s {
		if el != "" {
			calories, err := strconv.Atoi(el)
			if err != nil {
				log.Fatal(err)
			}

			t += calories
		}

		if el == "" {
			c = append(c, t)
			t = 0
		}
	}

	sort.Ints(c)

	m := sum(c[len(c)-3:])

	fmt.Println("The top three total calories sum amount is", m)
}
