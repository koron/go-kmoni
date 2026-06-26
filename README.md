# kmoni - Kyoushin Monitor

[![PkgGoDev](https://pkg.go.dev/badge/github.com/koron/kmoni)](https://pkg.go.dev/github.com/koron/kmoni)
[![Actions/Go](https://github.com/koron/kmoni/workflows/Go/badge.svg)](https://github.com/koron/kmoni/actions?query=workflow%3AGo)
[![Go Report Card](https://goreportcard.com/badge/github.com/koron/kmoni)](https://goreportcard.com/report/github.com/koron/kmoni)
[![Ask DeepWiki](https://deepwiki.com/badge.svg)](https://deepwiki.com/koron/kmoni)

Kyoushin Monitor ([強震モニタ](http://www.kmoni.bosai.go.jp/)) is one of
earthquake monitors in Japan.  When it detects an earthquake, then reports data
of that.

## How to use

Install `kmoni` command

    $ go install github.com/koron/go-kmoni/cmd/kmoni@latest

Run `kmoni` command and wait for earthquakes in Japan.

    $ kmoni

When there are some earthquakes, you'll get this like log.

    2015/12/02 01:08:08 茨城県南部 最終報 M4.2 深さ50km 最大予想震度3
