package handler

import (
	"net/http"

	"github.com/jinmatt/twtrgo"
	"github.com/jinmatt/twtrgo/errors"
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
// so can be swapped out with any other prefered implementations
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
		template.RenderError(err, w)
		return
	}

	template.RenderHome(tweets, w)
}

// handleSearch http handler for the app's search tweets feed
func (h *TweetHandler) handleSearch(w http.ResponseWriter, r *http.Request) {
	// get searched keyword from query string
	keyword := r.URL.Query().Get("q")
	if keyword == "" {
		w.WriteHeader(http.StatusBadRequest)
		template.RenderError(errors.ErrNoKeyword, w)
		return
	}

	tweets, err := h.TweetService.Search(keyword)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		template.RenderError(err, w)
		return
	}

	template.RenderSearch(keyword, tweets, w)
}
