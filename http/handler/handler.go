package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// Handler type to hold all http handlers to services
type Handler struct {
	TweetHandler *TweetHandler
}

// NewHandler inits all handlers
func NewHandler() *Handler {
	return &Handler{
		TweetHandler: NewTweetHandler(),
	}
}

// MakeHandler defines routes and inits a basic middleware
func (h *Handler) MakeHandler() (http.Handler, error) {
	router := mux.NewRouter()

	// ---------
	//  Define routes here
	// ---------
	router.HandleFunc("/", h.TweetHandler.handleHomeFeed)
	router.HandleFunc("/search", h.TweetHandler.handleSearch)

	// middleware for some basic routes access logs
	n := negroni.Classic()
	n.UseHandler(router)

	return n, nil
}
