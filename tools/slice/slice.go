package slice

import (
	"math/rand"
	"time"
)

func ShuffleArray[T comparable](arr []T) []T {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	r.Shuffle(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] })

	return arr
}
