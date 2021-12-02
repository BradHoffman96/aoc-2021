package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var sum, lastSum int

func main() {

	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	s := bufio.NewScanner(f)

	var q []int
	var ctr int
	for s.Scan() {
		sum = 0

		depth, err := strconv.Atoi(s.Text())
		if err != nil {
			log.Fatal(err)
		}

		q = append(q, depth)

		if len(q) == 3 {
			for _, v := range q {
				sum = sum + v
			}

			if lastSum != 0 && lastSum < sum {
				ctr++
			}

			lastSum = sum
			q = q[1:]
		}
	}

	fmt.Println(ctr)
}
