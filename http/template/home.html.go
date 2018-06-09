// Code generated by hero.
// source: /Users/jinmatt/go/src/github.com/jinmatt/twtrgo/http/template/src/home.html
// DO NOT EDIT!
package template

import (
	"io"

	"github.com/jinmatt/twtrgo"
	"github.com/shiyanhui/hero"
)

func RenderHome(tweets []*twtrgo.Tweet, w io.Writer) {
	_buffer := hero.GetBuffer()
	defer hero.PutBuffer(_buffer)
	_buffer.WriteString(`<!doctype html>
<html lang="en">
  <head>
    <!-- Required meta tags -->
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">

    <!-- Bootstrap CSS -->
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.1/css/bootstrap.min.css" integrity="sha384-WskhaSGFgHYWDcbwN70/dfYBj47jz9qbsMId/iRN3ewGhXQFZCSftd1LZCfmhktB" crossorigin="anonymous">

    <title>TwtrGo! - A badly named twitter app!</title>
  </head>
  <body>
    <div class="container">
      <div class="row align-items-start">
        <div class="col align-self-center">
          `)
	_buffer.WriteString(`<!-- navbar template -->
<nav class="navbar sticky-top navbar-dark my-3 shadow rounded" style="background-color: #794bc4;">
  <!-- brand -->
  <a class="navbar-brand" href="/">TwtrGo</a>

  <!-- search form -->
  <form class="form-inline" method="get" action="/search">
    <input class="form-control mr-sm-2" type="search" placeholder="Search" aria-label="Search" name="q">
    <button class="btn btn-outline-light my-2 my-sm-0" type="submit">Search</button>
  </form>
</nav>
`)
	_buffer.WriteString(`

    <div class="row">
      <div class="col">
        `)
	for _, tweet := range tweets {
		_buffer.WriteString(`<!-- Tweet list item -->
<div class="media my-3 rounded shadow-sm" style="border: 1px solid #fee2e6!important;">
  <img class="align-self-start ml-3 mt-3 rounded-circle" src="`)
		hero.EscapeHTML(tweet.User.ProfileImageURL, _buffer)
		_buffer.WriteString(`">
  <div class="media-body">
    <div class="card border-0">
      <div class="card-body p-3">
        <h6 class="card-title mb-2">`)
		hero.EscapeHTML(tweet.User.Name, _buffer)
		_buffer.WriteString(`
          <small class="text-muted">@`)
		hero.EscapeHTML(tweet.User.ScreenName, _buffer)
		_buffer.WriteString(`  &middot; `)
		hero.EscapeHTML(tweet.CreatedAt, _buffer)
		_buffer.WriteString(`</small>
        </h6>
        <p class="card-text font-weight-light">@`)
		hero.EscapeHTML(tweet.Status, _buffer)
		_buffer.WriteString(`</p>
      </div>
    </div>
  </div>
</div>
`)
	}
	if len(tweets) == 0 {
		_buffer.WriteString(`<!-- Tweet list item -->
<div class="media my-3 rounded shadow-sm" style="border: 1px solid #fee2e6!important;">
  <div class="media-body">
    <div class="card border-0">
      <div class="card-body p-3">
        <p class="card-text font-weight-light">No tweets were found!</p>
      </div>
    </div>
  </div>
</div>
`)
	}
	_buffer.WriteString(`
      </div>
    </div>

`)

	_buffer.WriteString(`
        </div>
      </div>
    </div>
  </body>
</html>
`)
	w.Write(_buffer.Bytes())

}