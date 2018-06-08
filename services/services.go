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

var (
	twitterApi *anaconda.TwitterApi
)

func Init(config *config.Config) error {
	initOnce.Do(func() {
		// initialize twitter api client
		anaconda.SetConsumerKey(config.TwitterConsumerKey)
		anaconda.SetConsumerSecret(config.TwitterConsumerSecret)
		twitterApi = anaconda.NewTwitterApi(config.TwitterAccessToken, config.TwitterAccessTokenSecret)
	})

	return nil
}

func Close() {
	closeOnce.Do(func() {
		twitterApi.Close()
	})
}

func TwitterAPI() *anaconda.TwitterApi {
	return twitterApi
}
