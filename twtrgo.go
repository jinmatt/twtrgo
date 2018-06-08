package twtrgo

type Tweet struct {
	ID     int64  `json:"id"`
	Status string `json:"status"`
}

type TweetService interface {
	Tweets() (tweets []*Tweet, err error)
	//SearchTweets([]*Tweet, error)
}
