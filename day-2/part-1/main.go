package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var depth, dist int

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	s := bufio.NewScanner(f)

	for s.Scan() {
		cmd := s.Text()
		args := strings.Split(cmd, " ")

		instr := args[0]
		val, err := strconv.Atoi(args[1])
		if err != nil {
			log.Fatal(err)
		}

		switch instr {
		case "forward":
			dist = dist + val
		case "down":
			depth = depth + val
		case "up":
			depth = depth - val
		default:
			log.Fatal("unrecognized instruction:", cmd)
		}
	}

	fmt.Println("Distance * Depth:", dist*depth)
}
