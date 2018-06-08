package twitter

import (
	"net/url"

	"github.com/ChimeraCoder/anaconda"
	"github.com/jinmatt/twtrgo"
)

// Check to see whether `TweetService` implements `twtrgo.TweetService`
var _ twtrgo.TweetService = &TweetService{}

// TweetService implements `twtrgo.TweetService`
// `Api` type can be swapped out to use any Twitter client
// or custom http implementions with the `services` package
type TweetService struct {
	Api *anaconda.TwitterApi
}

// HomeFeed gets the home timeline feeds from twitter
func (t *TweetService) HomeFeed() (tweets []*twtrgo.Tweet, err error) {
	timeline, err := t.Api.GetHomeTimeline(url.Values{
		"count":            []string{"50"},
		"exclude_replies":  []string{"true"},
		"include_entities": []string{"false"},
	})
	if err != nil {
		return nil, err
	}

	for _, t := range timeline {
		tweet := &twtrgo.Tweet{
			ID:     t.Id,
			Status: t.Text,
			User: &twtrgo.User{
				ID:              t.User.Id,
				Name:            t.User.Name,
				ScreenName:      t.User.ScreenName,
				ProfileImageURL: t.User.ProfileImageUrlHttps,
			},
		}

		tweets = append(tweets, tweet)
	}

	return tweets, nil
}
