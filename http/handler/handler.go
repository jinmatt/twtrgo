package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

type Handler struct {
	TweetHandler *TweetHandler
}

func NewHandler() *Handler {
	return &Handler{
		TweetHandler: NewTweetHandler(),
	}
}

func (h *Handler) MakeHandler() (http.Handler, error) {
	router := mux.NewRouter()

	// define routes here
	router.HandleFunc("/", h.TweetHandler.handleGetTweets)

	// for some basic routes access logs
	n := negroni.Classic()
	n.UseHandler(router)

	return n, nil
}
