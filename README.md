# everyman-rss

RSS feeds for Everyman Cinema film releases.

## Why?

I wanted a feed of new film releases within my
[Inoreader](https://www.inoreader.com) client.

![Screenshot of Inoreader](/assets/inoreader-screenshot.png)

## Feeds

Each feed is valid RSS served over HTTPS.

### Firehose Feeds

The "firehose" feeds contain all new films as they appear on Everyman Cinema:

```
https://everyman-rss.vercel.app/films
```

## Development

### Local

Run all of the endpoints locally using (see [`main.go`](/main.go)):

```
make run
```

### Caching

- Feeds are cached for 5 minutes
- After that the stale feed will continue to be served, whilst simultaneously
  updating the cache in the background
- See: [Stale-While-Revalidate](https://vercel.com/docs/concepts/edge-network/caching#stale-while-revalidate)

## Credits

- Thanks to Everyman Cinema for providing the REST API
- Big fan of [Hacker News RSS](https://github.com/hnrss/hnrss) by
  [@edavis](https://github.com/edavis)
