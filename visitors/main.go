package main

import "fmt"

type visitor struct {
	Arrival   int
	Departure int
}

func main() {
	visitors := []visitor{
		{1, 4},
		{2, 5},
		{9, 12},
		{5, 9},
		{5, 12},
	}

	cnt := maxNumberOfVisitors(visitors, 5, 9)
	fmt.Printf("%d\n", cnt)
}

func maxNumberOfVisitors(visitors []visitor, start, end int) int {
	cnt := 0
	for _, v := range visitors {
		if v.Departure < v.Arrival {
			continue
		}

		if v.Departure <= start {
			continue
		}

		if v.Arrival <= end {
			cnt++
		} else if v.Arrival > start {
			cnt++
		}
	}

	return cnt
}
