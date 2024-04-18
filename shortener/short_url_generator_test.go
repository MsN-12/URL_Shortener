package shortener

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"log"
	"math/rand"
	"strings"
	"testing"
)

func TestGenerateShortURL(t *testing.T) {
	originalURL1 := "https://www.google.com/1"
	shortURL1 := GenerateShortURL(originalURL1)

	originalURL2 := "https://www.google.com/2"
	shortURL2 := GenerateShortURL(originalURL2)

	require.Equal(t, len(shortURL1), 8)
	require.NotEqual(t, shortURL1, shortURL2)
}

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func TestCollisionRate(t *testing.T) {
	urlMap := make(map[string]bool)
	collisions := 0
	count := 0

	for i := 0; i < 1000; i++ {
		s := RandomString(100)
		shortURL := GenerateShortURL(s)
		if urlMap[shortURL] {
			collisions++
		} else {
			urlMap[shortURL] = true
			count++
		}
	}

	assert.Equal(t, collisions, 0)
	log.Println("Collisions: ", collisions)
}
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}
