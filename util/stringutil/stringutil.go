// MIT License
// 
// Copyright (c) 2017 Nobi Hoang <nobi.hoa@gmail.com>
// 

package stringutil

import (
	"math/rand"
    "time"
)

// CHARSET alphanumric constant
const CHARSET = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// RandString returns a random string
// 
// References: https://www.calhoun.io/5-tips-for-using-strings-in-go-2
func RandString(length int) string {
	source := rand.NewSource(time.Now().UnixNano())
    b := make([]byte, length)
    for i := range b {
        b[i] = CHARSET[source.Int63()%int64(len(CHARSET))]
    }
    return string(b)
}

// Reverse returns the string in the reversed order
//
// https://github.com/golang/example/blob/master/stringutil/reverse.go#L21
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}