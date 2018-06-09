package twitter

import (
	"net/url"
	"time"

	"github.com/ChimeraCoder/anaconda"
	"github.com/dustin/go-humanize"
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

	tweets, err = parseTweets(timeline)
	if err != nil {
		return nil, err
	}

	return tweets, nil
}

// Search gets tweets based on a searched keyword
func (t *TweetService) Search(keyword string) (tweets []*twtrgo.Tweet, err error) {
	results, err := t.Api.GetSearch(keyword, url.Values{
		"count":            []string{"50"},
		"lang":             []string{"en"},
		"locale":           []string{"en"},
		"include_entities": []string{"false"},
	})
	if err != nil {
		return nil, err
	}

	tweets, err = parseTweets(results.Statuses)
	if err != nil {
		return nil, err
	}

	return tweets, nil
}

// parseTweets converts the list of tweets statuses to type `[]*twtrgo.Tweet`
func parseTweets(statuses []anaconda.Tweet) (tweets []*twtrgo.Tweet, err error) {
	for _, status := range statuses {
		createdAt, err := time.Parse(time.RubyDate, status.CreatedAt)
		if err != nil {
			return nil, err
		}
		tweet := &twtrgo.Tweet{
			ID:        status.Id,
			Status:    status.Text,
			CreatedAt: humanize.Time(createdAt),
			User: &twtrgo.User{
				ID:              status.User.Id,
				Name:            status.User.Name,
				ScreenName:      status.User.ScreenName,
				ProfileImageURL: status.User.ProfileImageUrlHttps,
			},
		}

		tweets = append(tweets, tweet)
	}

	return tweets, nil
}
