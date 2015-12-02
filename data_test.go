package kmoni

import (
	"encoding/json"
	"testing"
)

func TestLoadJSON(t *testing.T) {
	var d Data
	err := json.Unmarshal([]byte(`{
  "result":{
    "status":"success",
    "message": "",
    "is_auth":true
  },
  "report_time":"2015/12/02 01:08:08",
  "region_code":"",
  "request_time":"20151202010958",
  "region_name":"茨城県南部",
  "longitude":"140.1",
  "is_cancel":false,
  "depth":"50km",
  "calcintensity":"3",
  "is_final":true,
  "is_training":false,
  "latitude":"36.1",
  "origin_time":"20151202010715",
  "security":{
    "realm":"/kyoshin_monitor/static/jsondata/eew_est/",
    "hash":"b61e4d95a8c42e004665825c098a6de4"
  },
  "magunitude":"4.2",
  "report_num":"4",
  "request_hypo_type":"eew",
  "report_id":"20151202010723",
  "alertflg":"予報"
}`), &d)
	if err != nil {
		t.Fatal(err)
	}
	if !d.IsFinal {
		t.Fatal("IsFinal should be true:", d.IsFinal)
	}
}

func TestLoadEmptyJSON(t *testing.T) {
	var d Data
	err := json.Unmarshal([]byte(`{
  "result": {
    "status": "success",
    "message": "データがありません",
    "is_auth": true
  },
  "report_time": "",
  "region_code": "",
  "request_time": "20151202011500",
  "region_name": "",
  "longitude": "",
  "is_cancel": "",
  "depth": "",
  "calcintensity": "",
  "is_final": "",
  "is_training": "",
  "latitude": "",
  "origin_time": "",
  "security": {
    "realm": "/webservice/hypo/eew/",
    "hash": "5ca8b8104e01ceef0f061ad597606cbd87b492db"
  },
  "magunitude": "",
  "report_num": "",
  "request_hypo_type": "eew",
  "report_id": ""
}`), &d)
	if err != nil {
		t.Fatal(err)
	}
	if d.IsFinal {
		t.Fatal("IsFinal should be false")
	}
}
