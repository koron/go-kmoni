package kmoni

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Now fetches current Data of kyoushin monitor.
func Now() (*Data, error) {
	return Fetch(time.Now())
}

// Fetch fetches Data for specified time from kyoushin monitor.
func Fetch(t time.Time) (*Data, error) {
	// get JSON
	resp, err := http.Get(url("eew", t))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("kmoni server failure: %d", resp.StatusCode)
	}
	// parse JSON
	var d Data
	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return nil, err
	}
	return &d, nil
}

func jst() *time.Location {
	const loc = "Asia/Tokyo"
	l, err := time.LoadLocation(loc)
	if err != nil {
		l = time.FixedZone(loc, 9*60*60)
	}
	return l
}

func url(hypo string, t time.Time) string {
	s := t.In(jst()).Format("20060102150405")
	return fmt.Sprintf("http://www.kmoni.bosai.go.jp/webservice/hypo/%s/%s.json", hypo, s)
}
