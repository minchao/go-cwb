# go-cwb

[![Build Status](https://travis-ci.org/minchao/go-cwb.svg?branch=master)](https://travis-ci.org/minchao/go-cwb)
[![Go Report Card](https://goreportcard.com/badge/github.com/minchao/go-cwb)](https://goreportcard.com/report/github.com/minchao/go-cwb)
[![codecov](https://codecov.io/gh/minchao/go-cwb/branch/master/graph/badge.svg)](https://codecov.io/gh/minchao/go-cwb)

An unofficial Taiwan [Central Weather Bureau RESTful API](http://opendata.cwb.gov.tw/) library for Go.

This package is inspired by [go-github](https://github.com/google/go-github).

## Installation

```go
go get github.com/minchao/go-cwb
```

## Usage

```go
import "github.com/minchao/go-cwb/cwb"
```

You will need an account on [cwb.gov.tw](http://www.cwb.gov.tw/) to get an API key.

Construct a new CWB client, then use to access the CWB API.
For example:

```go
client := cwb.NewClient("APIKEY", nil)
```

Get the 36 hour weather forecasts:

```go
forecast, _, err := client.Forecasts.Get36HourWeather(context.Background(), nil, nil)
if err != nil {
    fmt.Println(err)
}
```

Get the 2 day townships weather forecasts by specific country:

```go
forecast, _, err := client.Forecasts.GetTownshipsWeatherByDataId(context.Background(), cwb.FTW2DayTaipeiCity, nil, nil)
```

Get the 7 day township weather forecasts by specific locations and elements:

```go
forecast, _, err := client.Forecasts.GetTownshipsWeatherByDataId(context.Background(),
    cwb.FTW7DayChiayiCity,
    []string{"阿里山鄉"},
    []string{"MinT", "MaxT", "WeatherDescription"})
```

## Implemented APIs

* 一般天氣預報-今明36小時天氣預報 (F-C0032-001)
* 鄉鎮天氣預報-單一鄉鎮市區預報資料 (F-D0047-001 - F-D0047-091)
* 鄉鎮天氣預報-全臺灣各鄉鎮市區預報資料 (F-D0047-093) 
* 自動氣象站-氣象觀測資料 (O-A0001-001)
* 自動雨量站-雨量觀測資料 (O-A0002-001)

[CWB Open data API documentation Version 1.2](http://opendata.cwb.gov.tw/opendatadoc/CWB_Opendata_API_V1.2.pdf)

## License

See the [LICENSE](LICENSE.md) file for license rights and limitations (MIT).
