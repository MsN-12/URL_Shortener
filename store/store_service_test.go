package store

import (
	"github.com/stretchr/testify/require"
	"testing"
)

var testStoreService = InitializeStore()

func TestInsertAndRetrieve(t *testing.T) {
	longURL := "https://www.exmple.com/sdafa/asfdasf/asfdas/sadfas/dsafa/asdf/asdf/asdf"
	shortURL := "XsDfa534sfs"

	SaveUrl(shortURL, longURL)
	retrivedURL, err := RetrieveUrl(shortURL)
	require.NoError(t, err)
	require.True(t, testStoreService.redisClient != nil)
	require.Equal(t, retrivedURL, longURL)
}
