# Everyman RSS

[![GoDoc](https://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://pkg.go.dev/github.com/revett/everyman-rss)
[![Builds](https://img.shields.io/github/checks-status/revett/everyman-rss/main?label=build&style=flat-square)](https://github.com/revett/everyman-rss/actions?query=branch%3Amain)
[![Go Report Card](https://goreportcard.com/badge/github.com/revett/everyman-rss?style=flat-square)](https://goreportcard.com/report/github.com/revett/everyman-rss)
[![Codecov](https://img.shields.io/codecov/c/github/revett/everyman-rss.svg?style=flat-square)](https://codecov.io/gh/revett/everyman-rss)
[![License](https://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://github.com/revett/everyman-rss/blob/main/LICENSE)

RSS feeds for [Everyman Cinema](https://www.everymancinema.com) film releases.

## Why?

I wanted a feed of new film releases within my
[Inoreader](https://www.inoreader.com) client for my local cinema.

![Screenshot of Inoreader](https://github.com/revett/everyman-rss/blob/main/assets/inoreader-screenshot.png?raw=true)

## Feeds

Each feed is an [RSS 2.0](https://www.rssboard.org/rss-specification) feed,
served over HTTPS.

### Cinema Feeds

You can get a feed of all new films for a specific Everyman Cinema:

```
https://everyman-rss.revcd.com/films?cinema=broadgate
```

See a list of all available cinemas
[below](https://everyman-rss.revcd.com/#cinemas).

### Caching

Feeds are cached for 5 minutes.
