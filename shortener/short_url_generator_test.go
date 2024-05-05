package shortener

import (
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
	count := 1000000
	for i := 0; i < count; i++ {
		s := RandomString(50)
		shortURL := GenerateShortURL(s)
		if urlMap[shortURL] {
			collisions++
		} else {
			urlMap[shortURL] = true
		}
	}
	collisionsRate := float64(collisions) / float64(count)
	log.Println("Collisions: ", collisions)
	log.Println("Collisions rate: ", collisionsRate)
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
