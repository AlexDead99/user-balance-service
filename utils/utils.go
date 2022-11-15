package utils

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func generateRandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func generateRandomName(size int) string {
	var sb strings.Builder

	for i := 0; i < size; i++ {
		c := alphabet[rand.Intn(len(alphabet))]
		sb.WriteByte(c)
	}

	return sb.String()
}

func CreateOwner() string {
	return generateRandomName(6)
}

func CreateBalance() float32 {
	return float32(generateRandomInt(300, 5000))
}
