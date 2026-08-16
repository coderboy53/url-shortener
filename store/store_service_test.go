package store

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testStoreService = &StorageService{}

// initialize the store service for the tests
func init() {
	testStoreService = InitializeStore()
}


// asserts whether the provided condition holds true or not. Fails the test if otherwise
func TestStoreInit(t *testing.T) {
	assert.True(t, testStoreService.redisClient != nil)
}

func TestInsertionAndRetrieval(t *testing.T) {
	// test storage
	initialLink := "https://www.guru3d.com/news-story/spotted-ryzen-threadripper-pro-3995wx-processor-with-8-channel-ddr4,2.html"
	userUUId := "e0dba740-fc4b-4977-872c-d360239e6b1a"
	shortURL := "Jsz4k57oAX"
	StoreMapping(shortURL, initialLink, userUUId)
	// test retrieval
	retrievedUrl := RetrieveInitialUrl(shortURL)
	assert.True(t, initialLink == retrievedUrl)
}

