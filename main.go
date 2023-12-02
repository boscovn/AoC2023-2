package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const redAmount = 12
	const greenAmount = 13
	const blueAmount = 14
	sum := 0

	scanLoop:
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		s := strings.Split(text, ": ")
		index := strings.Split(s[0], " ")
		num, err := strconv.Atoi(index[1])
		if err != nil {
			break
		}
		games := strings.Split(s[1], "; ")
		for _, game := range(games){
			balls := strings.Split(game, ", ")
			if len(balls) <= 3 {
				m:= make(map[string]int)
				for _, ball := range(balls) {
					amountBall := strings.Split(ball, " ")
					amount, err := strconv.Atoi(amountBall[0])
					if err != nil {
						continue scanLoop
					}
					m[amountBall[1]] = amount
				}
				if m["red"] > redAmount || m["green"] > greenAmount || m["blue"] > blueAmount {
					continue scanLoop
				}
			} else {
				continue scanLoop
			}
		}
		sum += num


	}

	fmt.Println(sum)
}
