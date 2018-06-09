# twtrgo

A basic twitter app to search tweets and to display user's home feed.

See demo [here](https://twtrgo.herokuapp.com).

### Dependencies

* `dep` - Go dependencies management tool [https://golang.github.io/dep](https://golang.github.io/dep)
* `hero` - Html templating engine for Go [https://github.com/shiyanhui/hero](https://github.com/shiyanhui/hero)
  * Install `hero` cli `go get -u github.com/shiyanhui/hero/hero`

### How to

> Instructions are based on considering the repo is cloned into a [Go workspace](https://golang.org/doc/code.html#Workspaces)

1. Rename `dev.env` as `.env`, set exports for environment variables for local development, including [Twitter app credentials](https://developer.twitter.com/en/docs/basics/authentication/guides/access-tokens.html):
```
TWTRGO_ENV=default
PORT=8080
TWITTER_CONSUMER_KEY=<consumer-key>
TWITTER_CONSUMER_SECRET=<consumer-secret>
TWITTER_ACCESS_TOKEN=<access-token>
TWITTER_ACCESS_TOKEN_SECRET=<access-token-secret>
```
> Otherwise _export_ environment variables directly in terminal, whichever is preferable

2. Install `hero` cli as mentioned in **Dependencies**. Html templates are compiled into directory `http/templates`. If need to recompile templates from `http/templates/src` run:
```
$ make template
```
> Default package name will be `github.com/jinmatt/twtrgo/http/template`. The render functions can be found under _http/template/src_ inside _.html_ files(home/search/error)

3. Build app:
```
$ make build
```

4. Run tests:
```
$ make test
```

5. Run app:
```
$ make
```
> App starts on default port set in .env file _(PORT=8080 ; http://localhost:8080)_

### Package Layout

The app package layout is based on an approach to scale(Ref. [here](https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1)).

* `twtrgo` - The main package holds _Tweets_ type and _TweetService_ interface services should implement based on the requirements
* `twitter` - implements _twtrgo.TweetService_ interface with [Twitter API client](https://github.com/ChimeraCoder/anaconda). Can be swapped out with other implementations or to use a cache/db service
* `http/handler` - Handles http routes
* `services` - Handles global connection/client objects like API/DB/cache objects, so they are only initialized once and destroyed once
* `config` - Handles runtime configs
* `cmd/twtrgo` - App binary package, inits configs, services and starts http server
* `mock` - A mock package for _tests_, implements _twtrgo.TweetService_
* `test` - Test package, tests only http routes(behavioral tests)
