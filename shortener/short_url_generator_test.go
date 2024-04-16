package shortener

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// Todo : test for collision rate
func TestGenerateShortURL(t *testing.T) {
	originalURL1 := "https://www.google.com/2"
	shortURL1 := GenerateShortURL(originalURL1)

	originalURL2 := "https://www.google.com/2"
	shortURL2 := GenerateShortURL(originalURL2)

	require.Equal(t, len(shortURL1), 8)
	require.NotEqual(t, shortURL1, shortURL2)
}
