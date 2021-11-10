# everyman-rss

RSS feeds for Everyman Cinema film releases.

## Why?

I wanted a feed of new film releases within my
[Inoreader](https://www.inoreader.com) client.

## Usage

### Local

```
make run
```

### Vercel

> **Note**: see documentation on
> [Go runtime](https://vercel.com/docs/runtimes#official-runtimes/go).

Each [`http.HandlerFunc`](https://pkg.go.dev/net/http#HandlerFunc) within `api/`
is hosted as a separate serverless function.
