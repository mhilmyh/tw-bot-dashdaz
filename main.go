package main

import (
	"flag"
	"fmt"
)

func main()  {
	var tweet = flag.String("tweet", "", "Your tweet")
	flag.Parse()
	fmt.Printf("tweet: %s\n", *tweet)
}