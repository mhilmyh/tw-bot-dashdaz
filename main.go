package main

import (
	"dashbazbot/api"
	"dashbazbot/env"
)

func main()  {
	env.Load()
	q := api.GetQuote()
	api.CreateTweet(q.Content + " --" + q.Author)
}