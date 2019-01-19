package stringsutil

import (
	"fmt"
	"gotils/numbersutils"
	"math"
)

func dist(x1, y1, x2, y2 int) int {
	return int(math.Sqrt(math.Pow(float64(x2-x1), 2) + (math.Pow(float64(y2-y1), 2))))
}

// LevenshteinDistance Calculates Levenshtein distance between two strings
// n,m are the lengths of the strings respectively
func LevenshteinDistance(text1, text2 string) int {
	m := make([][]int, len(text1))
	for i := range text1 {
		m[i] = make([]int, len(text2))
		for j := range text2 {
			m[i][j] = -1
		}
	}

	return calculateLevenshteinDistance(text1, text2, len(text1)-1, len(text2)-1, m)
}

// LevenshteinDistanceWithPath Calculates levenshtein distance with path
func LevenshteinDistanceWithPath(text1, text2 string) (int, []string) {
	m := make([][]int, len(text1))
	x := make([]string, 0)
	for i := range text1 {
		m[i] = make([]int, len(text2))
		for j := range text2 {
			m[i][j] = -1
		}
	}

	return calculateLevenshteinDistanceWithPath(text1, text2, len(text1)-1, len(text2)-1, m, &x), x
}

func calculateLevenshteinDistanceWithPath(text1, text2 string, n, m int, store [][]int, stack *[]string) int {

	if n < 0 {
		return m + 1
	}
	if m < 0 {
		return n + 1
	}
	x := store[n][m]
	if x != -1 {
		return x
	}

	if text1[n] == text2[m] {
		store[n][m] = calculateLevenshteinDistanceWithPath(text1, text2, n-1, m-1, store, stack)
	} else {
		s1 := make([]string, 0)
		s2 := make([]string, 0)
		s3 := make([]string, 0)
		subs := calculateLevenshteinDistanceWithPath(text1, text2, n-1, m-1, store, &s1)
		dels := calculateLevenshteinDistanceWithPath(text1, text2, n, m-1, store, &s2)
		ins := calculateLevenshteinDistanceWithPath(text1, text2, n-1, m, store, &s3)
		min, v := numbersutils.MinInt(subs, dels, ins)
		// *stack = (*stack)[1:]
		if v == 0 {
			*stack = append(*stack, s1...)
			*stack = append(*stack, fmt.Sprintf("Subs %s with %s", string(text1[n-1]), string(text2[m-1])))
		} else if v == 1 {
			*stack = append(*stack, s2...)
			*stack = append(*stack, fmt.Sprintf("Del %s", string(text2[m])))
		} else if v == 2 {
			*stack = append(*stack, s3...)
			*stack = append(*stack, fmt.Sprintf("Ins %s", string(text2[m+1])))
		}
		store[n][m] = 1 + min
	}
	return store[n][m]
}

func calculateLevenshteinDistance(text1, text2 string, n, m int, store [][]int) int {

	if n < 0 {
		// store[n][m] = m + 1
		return m + 1
	}
	if m < 0 {
		// store[n][m] = n + 1
		return n + 1
	}
	x := store[n][m]
	if x != -1 {
		return x
	}

	if text1[n] == text2[m] {
		store[n][m] = calculateLevenshteinDistance(text1, text2, n-1, m-1, store)
	} else {
		subs := calculateLevenshteinDistance(text1, text2, n-1, m-1, store)
		dels := calculateLevenshteinDistance(text1, text2, n, m-1, store)
		ins := calculateLevenshteinDistance(text1, text2, n-1, m, store)
		min, _ := numbersutils.MinInt(subs, dels, ins)
		store[n][m] = 1 + min
	}
	return store[n][m]
}
