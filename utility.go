package main

import "math"

func toc(i int) uint8 {
	switch i {
	case 0:
		return 'E'
	case 1:
		return 'Y'
	case 2:
		return 'R'
	case 3:
		return 'G'
	case 4:
		return 'B'
	case 5:
		return 'W'
	}
	return math.MaxUint8
}

func toi(i uint8) int {
	switch i {
	case 'E':
		return 0
	case 'Y':
		return 1
	case 'R':
		return 2
	case 'G':
		return 3
	case 'B':
		return 4
	case 'W':
		return 5
	}
	return -1
}

func Swap(i1 *int, i2 *int) {
	tmp := *i1
	*i1 = *i2
	*i2 = tmp
}