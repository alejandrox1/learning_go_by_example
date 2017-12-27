package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"

	"github.com/google/go-github/github"
	"github.com/nlopes/slack"
)


func main() {
	port := 8081

	log.Println("Server started")
	http.HandleFunc("/webhook", handleWebhook)

	log.Printf("Server starting on port: %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}


func handleWebhook(w http.ResponseWriter, r *http.Request) {
	// read input into []byte buffer
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("error reading request body: err=%s\n", err)
		return
	}
	defer r.Body.Close()

	// paser webhook into an event (interface{} type)
	event, err := github.ParseWebHook(github.WebHookType(r), payload)
	if err != nil {
		log.Printf("could not parse webhook: err=%s\n", err)
		return
	}

	api := slack.New("x234326tahW8727S6d")
	devopsChannel := "7KFS7"
	var msg string

	log.Println("\n ----------------------------------------------------- ")
	//log.Printf("%+v\n", event)
	log.Printf("%v", reflect.TypeOf(event))

	switch e := event.(type) {
	case *github.PushEvent:
		// commit, push
		if e.Ref != nil {
			for _, commit := range e.Commits {
				msg := *commit.Author.Name + " -> \"" + *commit.Message + "\" on " + *e.Repo.Name + "/" + (*e.Ref)[len("refs/heads/"):]
				log.Printf("%s", msg)

				params := slack.PostMessageParameters{}
				_, _, err = api.PostMessage(devopsChannel, msg, params)
				if err != nil {
					fmt.Printf("%s\n", err)
					return
				}
			}
		}
	case *github.WatchEvent:
		//log.Printf("%v\n", *e.Action)
		// https://developer.github.com/v3/activity/events/types/#watchevent
		if e.Action != nil {
			msg = *e.Sender.Login + " starred repository " + *e.Repo.FullName
			log.Printf(msg)

			params := slack.PostMessageParameters{}
			_, _, err = api.PostMessage(devopsChannel, msg, params)
			if err != nil {
				fmt.Printf("%s\n", err)
				return
			}
		}
	default:
		log.Printf("unkown event type %s\n", github.WebHookType(r))
		return
	}
}
