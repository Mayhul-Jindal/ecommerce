package util

import (
    "fmt"
    "math/rand"
    "strings"
    "time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// RandomInt generates a random integer between min and max
func RandomInt(min, max int32) int32 {
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    return min + r.Int31n(max-min+1)
}

func RandomBigInt(min, max int64) int64 {
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    return min + r.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
    var sb strings.Builder
    k := len(alphabet)

    r := rand.New(rand.NewSource(time.Now().UnixNano()))

    for i := 0; i < n; i++ {
        c := alphabet[r.Intn(k)]
        sb.WriteByte(c)
    }

    return sb.String()
}

// RandomEmail generates a random email
func RandomEmail() string {
    return fmt.Sprintf("%s@email.com", RandomString(6))
}

func RandomTags() []string {
	var tagsArr []string
	n := int(RandomInt(1, 5))

	for i := 0; i < n; i++ {
		tagsArr = append(tagsArr, RandomString(4))
	}

	return tagsArr
}