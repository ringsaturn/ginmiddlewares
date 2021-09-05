# Gin middlewares for personal usage [![ci](https://github.com/ringsaturn/ginmiddlewares/actions/workflows/ci.yml/badge.svg)](https://github.com/ringsaturn/ginmiddlewares/actions/workflows/ci.yml)

Example: [`cmd/example/main.go`](cmd/example/main.go)

## Background

I just want to add `X-Response-Time` in the response, but there's no way to do that.

- <https://github.com/gin-gonic/gin/issues/1757>
- <https://github.com/gin-gonic/gin/issues/2493>

One day I found by accident that another project that used
<https://github.com/vearne/gin-timeout>, I can change header in a middleware
I defined without introducing many change in my code, so I dig into the timeout
repo and found the customed
[`TimeoutWriter `](https://github.com/vearne/gin-timeout/blob/v0.0.7/writer.go#L11-L23).

So I copy&paste it as a [middleware](xinject/inject.go) for gin.
