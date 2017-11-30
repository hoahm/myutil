// MIT License
// 
// Copyright (c) 2017 Nobi Hoang <nobi.hoa@gmail.com>
// 

package numutil

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// Xrand returns the random integer in the range of (min, max)
func Xrand(min, max int) int {
	rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}

// AtoSI split string by separator and convert to an Array of Integer
func AtoSI(s, sep string) []int {
	var a []int
	for _, c := range strings.Split(s, sep) {
		num, _ := strconv.Atoi(c)
		a = append(a, num)
	}
	return a
}

// IsDigit check whether a string is a decimal digit
func IsDigit(v string) bool {
	if _, err := strconv.Atoi(v); err == nil {
		return true
	}
	return false
}