package store

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

var redisAddr = os.Getenv("REDIS_ADDR")
var redisPassword = os.Getenv("REDIS_PASSWORD")
var testStoreService = InitializeStore(redisAddr, redisPassword)

func TestInsertAndRetrieve(t *testing.T) {
	longURL := "https://www.exmple.com/sdafa/asfdasf/asfdas/sadfas/dsafa/asdf/asdf/asdf"
	shortURL := "XsDfa534sfs"

	require.True(t, testStoreService.redisClient != nil)

	err := SaveUrl(shortURL, longURL)
	require.NoError(t, err)

	retrievedURL, err := RetrieveUrl(shortURL)
	require.NoError(t, err)
	require.Equal(t, retrievedURL, longURL)
}
