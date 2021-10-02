package api

import (
	"dashbazbot/utils"
	"log"
	"strconv"
)

func CreateTweet(s string)  {
	creds := utils.NewCreds()

	c, err  := creds.GetTwitterClient()
	if err != nil {
		log.Fatalln(err)
	}

	tw, _, err := c.Statuses.Update(s, nil)
	if err != nil {
		log.Fatalln(err)
	} 

	log.Println("[tweet:timestamp] " + tw.CreatedAt)
	log.Println("[tweet:id] " + strconv.Itoa(int(tw.ID)))
}