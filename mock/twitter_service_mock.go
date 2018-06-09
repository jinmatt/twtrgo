package mock

import "github.com/jinmatt/twtrgo"

// MockTweetService implements `twtrgo.TweetService`
type MockTweetService struct{}

func (t *MockTweetService) HomeFeed() (tweets []*twtrgo.Tweet, err error) {
	return tweets, nil
}

func (t *MockTweetService) Search(keyword string) (tweets []*twtrgo.Tweet, err error) {
	return tweets, nil
}
