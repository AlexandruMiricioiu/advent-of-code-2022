package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	s := strings.Split(string(content), "\n")

	m := 0
	t := 0
	for _, el := range s {
		if el != "" {
			calories, err := strconv.Atoi(el)
			if err != nil {
				log.Fatal(err)
			}

			t += calories
		}

		if el == "" {
			if t > m {
				m = t
			}

			t = 0
		}
	}

	fmt.Println("The maximum calories are", m)
}
