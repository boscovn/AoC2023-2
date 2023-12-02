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
	// const redAmount = 12
	// const greenAmount = 13
	// const blueAmount = 14
	sum := 0

scanLoop:
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		s := strings.Split(text, ": ")
		// index := strings.Split(s[0], " ")
		// num, err := strconv.Atoi(index[1])
		// if err != nil {
		// 	break
		// }
		games := strings.Split(s[1], "; ")
		m := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		for _, game := range games {
			balls := strings.Split(game, ", ")
			if len(balls) <= 3 {
				for _, ball := range balls {
					amountBall := strings.Split(ball, " ")
					amount, err := strconv.Atoi(amountBall[0])
					if err != nil {
						continue scanLoop
					}
					m[amountBall[1]] = max(amount, m[amountBall[1]])
				}
				// if m["red"] > redAmount || m["green"] > greenAmount || m["blue"] > blueAmount {
				// 	continue scanLoop
				// }
			} 
		}
		power := 1
		for k, v := range(m) {
			fmt.Println(k)
			if v >= 0 {
				power = power * v
			}
		}
		sum += power

	}

	fmt.Println(sum)
}
