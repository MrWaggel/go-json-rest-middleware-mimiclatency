# go-json-rest-middleware-mimiclatency
==================================

[![godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/MrWaggel/go-json-rest-middleware-mimiclatency) [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/MrWaggel/go-json-rest-middleware-mimiclatency/master/LICENSE)

Middleware to produce latency for [Go-Json-Rest](https://github.com/ant0ine/go-json-rest).

This is ideally used during development, to test how clients behave with high latency.

### Installation

You can do the usual:

    go get github.com/MrWaggel/go-json-rest-middleware-mimiclatency

### Usage

#### Import
Add the following imports.
```go
import (
	"github.com/MrWaggel/go-json-rest-middleware-mimiclatency"
	"time"
)
```
#### Add middleware
To add a delay of two seconds to your api requests.

```go
api.Use(&mimiclatency.MimicLatencyMiddleWare{
	Duration:time.Second * 2,
})
```
