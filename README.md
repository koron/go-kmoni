# kmoni - Kyoushin Monitor

Kyoushin Monitor ([強震モニタ](http://www.kmoni.bosai.go.jp/)) is one of
earthquake monitors in Japan.  When it detects an earthquake, then reports data
of that.

## How to use

Install `kmoni` command

    $ go install github.com/koron/go-kmoni/cmd/kmoni

Run `kmoni` command and wait for earthquakes in Japan.

    $ kmoni

When there are some earthquakes, you'll get this like log.

    2015/12/02 01:08:08 茨城県南部 最終報 M4.2 深さ50km 最大予想震度3
