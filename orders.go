package main

import (
	"fmt"
	"strings"
	"strconv"
)

func proceed(data string) int {
	split := strings.Split(data, " ")
	barberian, err := strconv.Atoi(split[0])
	if err != nil {
		fmt.Println("ERROR: ", err)
		return 0
	}

	timePeople := 0
	
	for i, item := range split {
		if(i == 0) {
			timePeople += 0
		} else {
			num, err := strconv.Atoi(item)
			if err != nil {
				fmt.Println("ERROR: ", err)
				return 0
			}
			timePeople += num
		}
	}

	res := timePeople / barberian

	return res
}

func main() {
	fmt.Println("RES: ", proceed("3 5 5 5 5 5"))
}