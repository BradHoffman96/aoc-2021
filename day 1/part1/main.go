package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	s := bufio.NewScanner(f)

	s.Scan()
	depth, err := strconv.Atoi(s.Text())
	if err != nil {
		log.Fatal(err)
	}

	var ctr int
	for s.Scan() {
		newDepth, err := strconv.Atoi(s.Text())
		if err != nil {
			log.Fatal(err)
		}

		if newDepth > depth {
			ctr++
		}

		depth = newDepth
	}

	fmt.Println(ctr)
}
