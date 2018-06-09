package services

import (
	"sync"

	"github.com/jinmatt/twtrgo/config"

	"github.com/ChimeraCoder/anaconda"
)

var (
	initOnce  sync.Once
	closeOnce sync.Once
)

// global variables to hold service objects
// db, cache or API connection objects can be added here
var (
	twitterApi *anaconda.TwitterApi
)

// Init helps in initializing external connections and APIs once through out the lifecycle of the http server
func Init(config *config.Config) error {
	initOnce.Do(func() {
		// initialize twitter api client
		anaconda.SetConsumerKey(config.TwitterConsumerKey)
		anaconda.SetConsumerSecret(config.TwitterConsumerSecret)
		twitterApi = anaconda.NewTwitterApi(config.TwitterAccessToken, config.TwitterAccessTokenSecret)
	})

	return nil
}

// Close cleans up connections opened by Init()
func Close() {
	closeOnce.Do(func() {
		twitterApi.Close()
	})
}

// TwitterAPI return the global variable with twitter API connection object
func TwitterAPI() *anaconda.TwitterApi {
	return twitterApi
}
