package utils

import (
	"math/rand"
)

// Shuffles a slice
func ShuffleSlice[T any](items []T) {
	rand.Shuffle(len(items), func(i, j int) {
		items[i], items[j] = items[j], items[i]
	})
}
