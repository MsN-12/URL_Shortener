package store

import (
	"github.com/stretchr/testify/require"
	"testing"
)

var testStoreService = &StorageService{}

func init() {
	testStoreService = InitializeStore()
}
func TestInsertAndRetrieve(t *testing.T) {
	longURL := "https://www.exmple.com/sdafa/asfdasf/asfdas/sadfas/dsafa/asdf/asdf/asdf"
	shortURL := "XsDfa534sfs"

	SaveUrl(shortURL, longURL)
	retrivedURL := RetriveUrl(shortURL)

	require.True(t, testStoreService.redisClient != nil)
	require.Equal(t, retrivedURL, longURL)
}
