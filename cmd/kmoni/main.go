package main

import (
	"fmt"
	"time"

	"github.com/koron/go-kmoni"
)

func data2str(d *kmoni.Data) string {
	if d.ReportID == "" {
		return ""
	}
	n := fmt.Sprintf("第%s報", d.ReportNum)
	if d.IsFinal {
		n = "最終報"
	}
	return fmt.Sprintf("%s %s %s M%s 深さ%s 最大予測震度%s", d.ReportTime, n,
		d.RegionName, d.Maginitude, d.Depth, d.Calcintensity)
}

func put(d *kmoni.Data) {
	s := data2str(d)
	if s == "" {
		return
	}
	fmt.Println(s)
}

func main() {
	for {
		d, err := kmoni.Now()
		if err != nil {
			panic(err)
		}
		put(d)
		time.Sleep(5 * time.Second)
	}
}
