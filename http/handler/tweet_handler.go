package handler

import (
	"net/http"

	"github.com/jinmatt/twtrgo"
	"github.com/jinmatt/twtrgo/http/template"
	"github.com/jinmatt/twtrgo/services"
	"github.com/jinmatt/twtrgo/twitter"
)

type TweetHandler struct {
	TweetService twtrgo.TweetService
}

func NewTweetHandler() *TweetHandler {
	return &TweetHandler{
		TweetService: &twitter.TweetService{
			Api: services.TwitterAPI(),
		},
	}
}

func (h *TweetHandler) handleGetTweets(w http.ResponseWriter, r *http.Request) {
	tweets, err := h.TweetService.Tweets()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	template.Feed(tweets, w)
}
