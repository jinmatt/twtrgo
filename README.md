# twtrgo

A basic twitter app to search tweets and to display user's home feed.

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

2. Install `hero` cli as mentioned in **Dependencies**. Html templates are compiled into directory `http/templates`. If need to recompile templates from `http/templates/src` run:
```
$ make template
```

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
