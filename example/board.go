package main

import (
	"fmt"
	"strconv"
)

const (
	sortedColor = "\033[1;36m%s\033[0m\t"
)

var pace int = 0 

var revealed = map[int]int{}

func board() {	
	headers := [10]int{ 1, 2, 16, 17, 31, 32, 46, 47, 61, 62 }
	for row := 0; row < 7; row++ {
		for col := 0; col < 10; col++ {
			_, ok := revealed[headers[col] + pace]
			if ok {
				fmt.Printf(sortedColor, strconv.Itoa(headers[col] + pace))
			} else {
				fmt.Print(headers[col] + pace, "\t")
			}
		}
		
		fmt.Println()
		pace += 2
	}
}

func main() {
	revealed[1] = 0
	revealed[3] = 0
	revealed[17] = 0
	revealed[21] = 0		
	revealed[69] = 0
	board()		
}	
