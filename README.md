# iGo - a Go interpreter, written in Go
[![Build Status](https://travis-ci.com/beeceej/iGo.svg?branch=master)](https://travis-ci.com/beeceej/iGo)
[![Go Report Card](https://goreportcard.com/badge/github.com/beeceej/iGo)](https://goreportcard.com/report/github.com/beeceej/iGo)

## Docker Build

```
TAG=latest make image
```

## Docker Usage
```
λ docker run --rm -d -p 9999:9999 beeceej/igo:latest
λ docker run --network host beeceej/igo:latest "igoclient 'func hi() string { return \"Hello\"}'"
```


## Supported

- Function Parsing
  - Single Function parsing
  - Multiple Function parsing at a time
- Expressions
  - Calling built in functions
  - Calling user defined functions

Follow the development of iGo here:


* [writing-a-go-interpreter-in-go](https://blog.beeceej.com/blog/writing-a-go-interpreter-in-go)
* [writing-a-go-interpreter-in-go-pt2](https://blog.beeceej.com/blog/writing-a-go-interpreter-in-go-pt2)
