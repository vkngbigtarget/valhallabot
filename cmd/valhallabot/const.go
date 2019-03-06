package main

import (
	"math/rand"
	"time"
)

const (
	UserImak = "238549436748398594>"
)

func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}
