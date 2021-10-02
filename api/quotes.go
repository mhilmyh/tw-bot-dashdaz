package api

import (
	"encoding/json"
	"log"
	"net/http"
)

const baseURLQuotes = "https://goquotes-api.herokuapp.com/api/v1/random?count=1"

type Quote struct {
	Content string
	Author string
}

func GetQuote() Quote {
	
	r, err := http.Get(baseURLQuotes)
	
	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	var m map[string] interface{}

	err = json.NewDecoder(r.Body).Decode(&m)

	if err != nil {
		panic(err)
	}

	q := m["quotes"].([]interface{})[0].(map[string]interface {})
	a := q["author"].(string)
	t := q["text"].(string)

	log.Println("[quotes:author] " + a)
	log.Println("[quotes:text] " + t)
	
	return Quote{
		Content: t,
		Author: a,
	}
}