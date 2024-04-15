package store

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

var testStoreService = &StorageService{}

func init() {
	testStoreService = InitializeStore()
}
func TestInsertAndRetrieve(t *testing.T) {
	longURL := "https://www.exmple.com/sdafa/asfdasf/asfdas/sadfas/dsafa/asdf/asdf/asdf"
	userID := uuid.New()
	shortURL := "XsDfa534sfs"

	SaveUrl(shortURL, longURL, userID.String())
	retrivedURL := RetriveUrl(shortURL, userID.String())

	require.True(t, testStoreService.redisClient != nil)
	require.Equal(t, retrivedURL, longURL)
}
