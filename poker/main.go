package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

type Rank []int

func (r Rank) Len() int {
	return len(r)
}

func (r Rank) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r Rank) Less(i, j int) bool {
	return r[i] > r[j]
}

func main() {
	ranks := map[string]int{
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
		"T": 10,
		"J": 11,
		"Q": 12,
		"K": 13,
		"A": 14,
	}

	s1 := score(ranks, os.Args[1:6])
	s2 := score(ranks, os.Args[6:11])
	fmt.Printf("1st deck: %v => %s\n", os.Args[1:6], s1)
	fmt.Printf("2nd deck: %v => %s\n", os.Args[6:11], s2)
	ret := strings.Compare(s1, s2)
	switch ret {
	case 0:
		fmt.Printf("tie\n")
	case 1:
		fmt.Printf("the winner is 1st deck\n")
	case -1:
		fmt.Printf("the winner is 2nd deck\n")
	}
}

func score(ranks map[string]int, ss []string) string {
	var suits []string
	var ranks_h []int
	var ranks_l []int
	is_flush := false
	is_straight := false
	var high_card int

	for _, v := range ss {
		x := strings.Split(v, "")
		suits = append(suits, x[1])

		if val, ok := ranks[x[0]]; ok {
			ranks_h = append(ranks_h, val)
			ranks_l = append(ranks_l, val)
		}
	}

	sort.Strings(suits)
	sort.Sort(Rank(ranks_h))
	sort.Sort(Rank(ranks_l))

	for ranks_l[0] == 14 {
		ranks_l[0] = 1
		sort.Sort(Rank(ranks_l))
	}

	if suits[0] == suits[4] {
		is_flush = true
	}

	if ranks_l[4] != 1 && ranks_l[0]-1 == ranks_l[1] &&
		ranks_l[1]-1 == ranks_l[2] &&
		ranks_l[2]-1 == ranks_l[3] &&
		ranks_l[3]-1 == ranks_l[4] {
		is_straight = true
		high_card = ranks_l[0]
	} else {
		high_card = ranks_h[0]
		if ranks_h[0]-1 == ranks_h[1] &&
			ranks_h[1]-1 == ranks_h[2] &&
			ranks_h[2]-1 == ranks_h[3] &&
			ranks_h[3]-1 == ranks_h[4] {
			is_straight = true
		}
	}

	if is_flush && is_straight {
		return _s(ranks_h, true, 9, high_card)
	} else if ranks_h[4] == ranks_h[1] {
		return _s(ranks_h, false, 7, 4, 0)
	} else if ranks_h[3] == ranks_h[0] {
		return _s(ranks_h, false, 7, 3, 4)
	} else if ranks_h[4] == ranks_h[3] && ranks_h[2] == ranks_h[0] {
		return _s(ranks_h, false, 6, 0, 4)
	} else if ranks_h[4] == ranks_h[2] && ranks_h[1] == ranks_h[0] {
		return _s(ranks_h, false, 6, 4, 0)
	} else if is_flush {
		return _s(ranks_h, false, 5, 0)
	} else if is_straight {
		return _s(ranks_h, true, 4, high_card)
	} else if ranks_h[4] == ranks_h[2] {
		return _s(ranks_h, false, 3, 4, 0, 1)
	} else if ranks_h[3] == ranks_h[1] {
		return _s(ranks_h, false, 3, 3, 0, 4)
	} else if ranks_h[2] == ranks_h[0] {
		return _s(ranks_h, false, 3, 2, 3, 4)
	} else if ranks_h[4] == ranks_h[3] && ranks_h[2] == ranks_h[1] {
		return _s(ranks_h, false, 2, 2, 4, 0)
	} else if ranks_h[4] == ranks_h[3] && ranks_h[1] == ranks_h[0] {
		return _s(ranks_h, false, 2, 1, 4, 2)
	} else if ranks_h[3] == ranks_h[2] && ranks_h[1] == ranks_h[0] {
		return _s(ranks_h, false, 2, 1, 3, 4)
	} else if ranks_h[4] == ranks_h[3] {
		return _s(ranks_h, false, 1, 4, 0, 1, 2)
	} else if ranks_h[3] == ranks_h[2] {
		return _s(ranks_h, false, 1, 3, 0, 1, 4)
	} else if ranks_h[2] == ranks_h[1] {
		return _s(ranks_h, false, 1, 2, 0, 3, 4)
	} else if ranks_h[1] == ranks_h[0] {
		return _s(ranks_h, false, 1, 1, 2, 3, 4)
	}
	return _s(ranks_h, false, 0, 0, 1, 2, 3, 4)
}

func _s(r []int, ref bool, a ...int) string {
	if !ref {
		for i := 1; i < len(a); i++ {
			v := a[i]
			a[i] = r[v]
		}
	}

	x := _conv(a)

	return strings.Join(x, "")
}

func _conv(i []int) []string {
	var r []string

	for _, v := range i {
		if v < 10 {
			r = append(r, fmt.Sprintf("0%d", v))
		} else {
			r = append(r, fmt.Sprintf("%d", v))
		}
	}

	return r
}
