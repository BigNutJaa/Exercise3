package main

import (
	"fmt"
	"sort"
	"strconv"
)

var wording string

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func main() {
	for {
		fmt.Println("Enter some word:")
		fmt.Scanln(&wording)
		if _, err := strconv.ParseFloat(wording, 64); err == nil {
			fmt.Printf("%q looks like a number.\n", wording)
		} else {
			break

		}
	}
	newWording := SortString(wording)
	fmt.Println("Input word =", wording)
	fmt.Println("New word(sort desc) =", newWording)

}
