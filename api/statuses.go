package api

import (
	"bytes"
	"encoding/base64"
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

	/**
	 * TODO: use oauth v1 with signature and consumer key
	 */
	b64 := base64.StdEncoding.EncodeToString([]byte(os.Getenv("EMAIL") + ":" + os.Getenv("PASS")))	
	req.Header.Add("Authentication", "Basic " + b64)
	req.Header.Add("Authorization", "OAuth " + os.Getenv("BEARER"))

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	var res map[string] interface {}

	json.NewDecoder(resp.Body).Decode(&res)

	fmt.Println(res)
}