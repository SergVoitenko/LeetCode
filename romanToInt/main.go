package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите число:")
	scanner.Scan()
	romanStr := scanner.Text()
	romanInt := romanToInt(romanStr)
	fmt.Printf("Integer :%d \n", romanInt)

}

func romanToInt(s string) int {
	var num int
	type dobleSim struct {
		sim string
		val int
	}
	lDubSim := []dobleSim{{sim: "CM", val: 900}, {sim: "CD", val: 400}, {sim: "XC", val: 90}, {sim: "XL", val: 40}, {sim: "IX", val: 9}, {sim: "IV", val: 4}}
	oneSim := map[string]int{"M": 1000, "D": 500, "C": 100, "L": 50, "X": 10, "V": 5, "I": 1}
	for _, simDub := range lDubSim {
		index := strings.Index(s, simDub.sim)
		if index >= 0 {
			num += simDub.val
			s = s[0:index] + s[index+2:]
		}

	}
	for i := 0; i <= len(s)-1; i++ {
		num += oneSim[string(s[i])]
	}
	return num

}
