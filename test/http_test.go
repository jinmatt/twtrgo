package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jinmatt/twtrgo/http/handler"
	"github.com/jinmatt/twtrgo/mock"
)

// TestHomeFeed tests home feed route
func TestHomeFeed(t *testing.T) {
	tweetHandler := &handler.TweetHandler{
		TweetService: &mock.MockTweetService{},
	}

	req, err := http.NewRequest("GET", "localhost:8080/", nil)
	if err != nil {
		t.Fatalf("Error creating http request: %s", err.Error())
	}
	rec := httptest.NewRecorder()

	tweetHandler.HandleHomeFeed(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("Expected status %d; got %d", http.StatusOK, res.StatusCode)
	}
}

// TestSearch tests search route /search
func TestSearch(t *testing.T) {
	tweetHandler := &handler.TweetHandler{
		TweetService: &mock.MockTweetService{},
	}

	// test cases
	tt := []struct {
		Name   string
		Query  string
		Status int
	}{
		{Name: "Found tweets", Query: "golang", Status: http.StatusOK},
		{Name: "No search keyword", Query: "", Status: http.StatusBadRequest},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "localhost:8080/search?q="+tc.Query, nil)
			if err != nil {
				t.Fatalf("Error creating http request: %s", err.Error())
			}
			rec := httptest.NewRecorder()

			tweetHandler.HandleSearch(rec, req)

			res := rec.Result()
			defer res.Body.Close()

			if res.StatusCode != tc.Status {
				t.Fatalf("Expected status %d; got %d", tc.Status, res.StatusCode)
			}
		})
	}
}
