package main

import "fmt"

var months = [...]string{1: "January",
	2:  "February",
	3:  "March",
	4:  "April",
	5:  "May",
	6:  "June",
	7:  "July",
	8:  "August",
	9:  "September",
	10: "October",
	11: "November",
	12: "December"}

func main() {
	Q2 := months[4:7]
	summer := months[6:9]

	fmt.Println(Q2)
	fmt.Println(summer)

	for _, q := range Q2 {
		for _, s := range summer {
			if q == s {
				fmt.Println(s, "appears in both")
			}
		}
	}
	endlessSummer := summer[:5]
	fmt.Println(endlessSummer)
}
