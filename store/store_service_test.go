package store

import (
	"github.com/stretchr/testify/require"
	"testing"
)

var testStoreService = InitializeStore()

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
