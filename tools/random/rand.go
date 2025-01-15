package random

import (
	"math/rand"
	"time"
)

func MakeRandomNumber(max int) int {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	return r.Intn(max)
}
