package utils

import (
	"log"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

type Creds struct {
	ConsumerKey					string
	ConsumerSecret 			string
	AccessToken					string
	AccessTokenSecret		string
}

func NewCreds() *Creds {
	return &Creds{
		ConsumerKey: os.Getenv("CONSUMER_KEY"),
		ConsumerSecret: os.Getenv("CONSUMER_SECRET"),
		AccessToken: os.Getenv("ACCESS_TOKEN"),
		AccessTokenSecret: os.Getenv("ACCESS_TOKEN_SECRET"),
	}
}

func (c *Creds) GetTwitterClient() (*twitter.Client, error) {
	conf := oauth1.NewConfig(c.ConsumerKey, c.ConsumerSecret)
	token := oauth1.NewToken(c.AccessToken, c.AccessTokenSecret)

	httpClient := conf.Client(oauth1.NoContext, token)
	twClient := twitter.NewClient(httpClient)

	verifParams := &twitter.AccountVerifyParams{
		SkipStatus: twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}

	user, _, err := twClient.Accounts.VerifyCredentials(verifParams)
	if err != nil {
		return nil, err
	}

	log.Println("[username] " + user.ScreenName)

	return twClient, nil
}