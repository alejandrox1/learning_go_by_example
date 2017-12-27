package main

import (
	//"encoding/json"
	"fmt"
	//"io"
	"io/ioutil"
	"log"
	"net/http"
	//"os"
	"reflect"

	"github.com/google/go-github/github"
)

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	//fmtPrintf("Headers: %v\n", r.Header)
	//_, err := io.Copy(os.Stdout, r.Body)
	//if err != nil {
	//    log.Println(err)
	//    return
	//}

	//webhookData := make(map[string]interface{})
	//err := json.NewDecoder(r.Body).Decode(&wenbhoodData)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	//fmt.Println("got webhook payload: ")
	//for k, v := range webhoodData {
	//	fmt.Printf("%s: %v\n", k, v)
	//}

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
	log.Println("\n ----------------------------------------------------- ")
	//log.Printf("%+v\n", event)
	log.Printf("%v", reflect.TypeOf(event))

	switch e := event.(type) {
	case *github.PushEvent:
		// commit, push
		if e.Ref != nil {
			for _, commit := range e.Commits {
				fmt.Printf("%s -> \"%s\" by %s\n", *e.Ref, *commit.Message, *commit.Author.Name)
			}
		}
	case *github.PullRequestEvent:
		// pull request
		if e.Action != nil {
			fmt.Println(*e.Action)
		}
	case *github.WatchEvent:
		//log.Printf("%v\n", *e.Action)
		// https://developer.github.com/v3/activity/events/types/#watchevent
		if e.Action != nil {
			log.Printf("%s -> %s starred repository %s\n", *e.Action, *e.Sender.Login, *e.Repo.FullName)
		}
	default:
		log.Printf("unkown event type %s\n", github.WebHookType(r))
		return
	}
}


func main() {
	log.Println("Server started")
	http.HandleFunc("/webhook", handleWebhook)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
