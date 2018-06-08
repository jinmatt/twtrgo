package twitter

import (
	"net/url"

	"github.com/ChimeraCoder/anaconda"
	"github.com/jinmatt/twtrgo"
)

var _ twtrgo.TweetService = &TweetService{}

type TweetService struct {
	Api *anaconda.TwitterApi
}

func (t *TweetService) Tweets() (tweets []*twtrgo.Tweet, err error) {
	res, err := t.Api.GetSearch("#4WordStoryOfMyLife", url.Values{
		"count":       []string{"50"},
		"result_type": []string{"mixed"},
		"lang":        []string{"en"},
	})
	if err != nil {
		return nil, err
	}

	for _, status := range res.Statuses {
		tweet := &twtrgo.Tweet{
			ID:     status.Id,
			Status: status.Text,
		}

		tweets = append(tweets, tweet)
	}

	return tweets, nil
}
