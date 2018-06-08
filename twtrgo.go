package twtrgo

// Tweet type to respresent Twitter tweet status
type Tweet struct {
	ID     int64  `json:"id"`
	Status string `json:"status"`
	User   *User  `json:"user,omitempty"`
}

// User type related to Tweet
type User struct {
	ID              int64  `json:"id"`
	Name            string `json:"name"`
	ScreenName      string `json:"screen_name"`
	ProfileImageURL string `json:"profile_image_url"`
}

// TweetService methods to be implemented by any service to fulfill the requirements
type TweetService interface {
	HomeFeed() (tweets []*Tweet, err error)
}
