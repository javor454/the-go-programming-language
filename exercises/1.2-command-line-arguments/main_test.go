package main

import (
	"math/rand"
	"strings"
	"testing"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var strs = generateStrings()

func stringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func randomString(length int) string {
	return stringWithCharset(length, charset)
}

func generateStrings() []string {
	sts := make([]string, 1000)
	for i := range sts {
		sts[i] = randomString(10)
	}
	return sts
}

func Join(strs []string) string {
	return strings.Join(strs, ",")
}

func BenchmarkJoin(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Join(strs)
	}
}

func Range(ss []string) string {
	var s string
	for _, v := range ss {
		s += " " + v
	}
	return s
}

func BenchmarkRange(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Range(strs)
	}
}

// RESULT:
// BenchmarkJoin-12     	1000000000	         0.0000180 ns/op
// BenchmarkRange-12    	1000000000	         0.0006656 ns/op

// RESULT2:
// BenchmarkJoin-12     	  213550	      5877 ns/op	   12288 B/op	       1 allocs/op
// BenchmarkRange-12    	    2490	    450874 ns/op	 5829023 B/op	    1000 allocs/op
