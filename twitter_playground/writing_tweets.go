package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/caarlos0/env"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

type twitterConfig struct {
	ConsumerKey string			`env:"CONSUMER_KEY,required"`
	ConsumerSecret string		`env:"CONSUMER_SECRET,required"`
	AccessToken string			`env:"ACCESS_TOKEN,required"`
	AccessTokenSecret string	`env:"ACCESS_TOKEN_SECRET,required"`
}

func main() {
	var tweetCount int
	var err error

	if len(os.Args) < 2 {
		tweetCount = 20
	} else {
		tweetCount, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Printf("%+v\n", err)
		}
	}

	err = writeTweets(tweetCount, "tweets.jsonl")
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
}

func writeTweets(tweetCount int, path string) error {
	// Get twitter credentials
	cfg := twitterConfig{}
	err := env.Parse(&cfg)
	if err != nil {
		return err
	}

	// Authentication for user
	config := oauth1.NewConfig(cfg.ConsumerKey, cfg.ConsumerSecret)
	token := oauth1.NewToken(cfg.AccessToken, cfg.AccessTokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// Home timeline
	tweets, _, err := client.Timelines.HomeTimeline(&twitter.HomeTimelineParams{
		Count: tweetCount,
	})
	if err != nil {
		return err
	}

	// Create file
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// write tweets to file
	w := bufio.NewWriter(file)
	for i, tweet := range tweets {
		fmt.Printf("%-2d) %-20s: %s\n", i, tweet.User.ScreenName, tweet.Text)

		jsonTweet, _ := json.Marshal(tweet)
		fmt.Fprintln(w, string(jsonTweet))
	}
	return w.Flush()
}

