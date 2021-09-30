package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const baseURLStatuses = "https://api.twitter.com/1.1/statuses/update.json"
func CreateTweet(s string)  {
	
	mapData := map[string]string {
		"status" : s,
	}

	reqData, err := json.Marshal(mapData)

	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest(http.MethodPost, baseURLStatuses, bytes.NewBuffer(reqData))

	if err != nil {
		panic(err)
	}

	req.Header.Add("oauth_consumer_key", os.Getenv("oauth_consumer_key"))
	req.Header.Add("oauth_nonce", "")
	req.Header.Add("oauth_signature", "")
	req.Header.Add("oauth_signature_method", "HMAC-SHA1")
	req.Header.Add("oauth_timestamp", "")
	req.Header.Add("oauth_token", "")
	req.Header.Add("oauth_version", "1.0")

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	var res map[string] interface {}

	json.NewDecoder(resp.Body).Decode(&res)

	fmt.Println(res)
}