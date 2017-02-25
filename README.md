# go-cwb

[![Go Report Card](https://goreportcard.com/badge/github.com/minchao/go-cwb)](https://goreportcard.com/report/github.com/minchao/go-cwb)
[![Build Status](https://travis-ci.org/minchao/go-cwb.svg?branch=master)](https://travis-ci.org/minchao/go-cwb)

An unofficial [CWB RESTful API](http://opendata.cwb.gov.tw/) library for Go.

This package is inspired by [go-github](https://github.com/google/go-github).

## Installation

```go
go get github.com/minchao/go-cwb
```

## Usage

```go
import "github.com/minchao/go-cwb/cwb"
```

Construct a new CWB client, then use to access the CWB API.

```go
client := cwb.NewClient("TOKEN", nil)
```

## License

See the [LICENSE](LICENSE.md) file for license rights and limitations (MIT).
