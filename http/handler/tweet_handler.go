package handler

import (
	"net/http"

	"github.com/jinmatt/twtrgo"
	"github.com/jinmatt/twtrgo/http/template"
	"github.com/jinmatt/twtrgo/services"
	"github.com/jinmatt/twtrgo/twitter"
)

// TweetHandler uses `twtrgo.TweetService`
// defines http handler methods for http routes
type TweetHandler struct {
	TweetService twtrgo.TweetService
}

// NewTweetHandler initializes TweetService
// `TweetService` is of type `twtrgo.TweetService`
// so can be swapped out with any other prefered implementions
func NewTweetHandler() *TweetHandler {
	return &TweetHandler{
		TweetService: &twitter.TweetService{
			Api: services.TwitterAPI(),
		},
	}
}

// handleHomeFeed http handler for the app's home page feed
func (h *TweetHandler) handleHomeFeed(w http.ResponseWriter, r *http.Request) {
	tweets, err := h.TweetService.HomeFeed()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	template.Render(tweets, w)
}
