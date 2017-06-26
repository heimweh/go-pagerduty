package integration

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	// charSetAlphaNum is the alphanumeric character set for use with
	// RandStringFromCharSet
	charSetAlphaNum = "abcdefghijklmnopqrstuvwxyz012346789"
)

func randResName() string {
	return fmt.Sprintf("test%s", randString(4))
}

// randString generates a random alphanumeric string of the length specified
func randString(strlen int) string {
	return randStringFromCharSet(strlen, charSetAlphaNum)
}

// randStringFromCharSet generates a random string by selecting characters from
// the charset provided
func randStringFromCharSet(strlen int, charSet string) string {
	reseed()
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = charSet[rand.Intn(len(charSet))]
	}
	return string(result)
}

// Seeds random with current timestamp
func reseed() {
	rand.Seed(time.Now().UTC().UnixNano())
}
