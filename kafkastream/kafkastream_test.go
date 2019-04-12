package kafkastream

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testInputJSON = `{
	"request": {
	  "client_ts": 1516711202,
	  "ref_screen": "Front",
	  "article_id": "18086074",
	  "position": 46,
	  "target": "derbund",
	  "type": "article_view",
	  "platform": "android"
	},
	"client_ts": 1516711202,
	"ts": 1516711203,
	"cookies": {},
	"ref_screen": "Front",
	"article_id": "18086074",
	"position": 46,
	"target": "derbund",
	"type": "article_view",
	"platform": "android",
	"requestHeaders": {
	  "tda-app-name": "Der Bund",
	  "tda-c1-creid": "1577680583332760801",
	  "tda-app-version": "7.3",
	  "tda-app-orientation": "portrait",
	  "tda-app-device": "phone",
	  "tda-app-resolution": "1080x1920",
	  "X-Forwarded-Host": "blackbeard.prod.tda.link",
	  "Timeout-Access": "<function1>",
	  "X-Original-URI": "/v1/event",
	  "tda-app-adid": "f62fadef-0809-42bb-9859-b59f7ca14bf6",
	  "tda-app-osversion": "7.0",
	  "X-Forwarded-For": "178.197.236.141",
	  "tda-uid": "ffffffff-f635-d114-0000-0000220b896c",
	  "Connection": "close",
	  "tda-geo-long": "7.4378486",
	  "X-Scheme": "https",
	  "X-Forwarded-Port": "443",
	  "Cache-Control": "max-age=0",
	  "tda-app-os": "android",
	  "bb_id": "ffffffff-f635-d114-0000-0000220b896c",
	  "X-Forwarded-Proto": "https",
	  "X-Real-Ip": "178.197.236.141",
	  "tda-geo-lat": "46.9487586",
	  "Accept-Encoding": "gzip",
	  "User-Agent": "okhttp/3.9.1",
	  "Host": "blackbeard.prod.tda.link"
	}
  }`
)

func TestJsonStructures(t *testing.T) {
	assert := assert.New(t)

	var parsedJSON ArticleWithUserID
	if err := json.Unmarshal([]byte(testInputJSON), &parsedJSON); err != nil {
		t.Fatal("JSON parsing failed ", err)
	}

	assert.Equal(int(parsedJSON.Timestamp), 1516711203)
	assert.Equal(parsedJSON.Request.ArticleID, "18086074")
	assert.Equal(parsedJSON.Request.Target, "derbund")
	assert.Equal(parsedJSON.RequestHeaders.UserID, "ffffffff-f635-d114-0000-0000220b896c")
}
