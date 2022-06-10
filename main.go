package main

import (
	"fmt"
	"net/http"

	"github.com/go-playground/webhooks/v6/github"
)

const (
	path = "/webhooks"
)

func main() {

	hook, _ := github.New(github.Options.Secret(""))

	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		payload, err := hook.Parse(r, github.ReleaseEvent, github.PullRequestEvent)
		if err != nil {
			if err == github.ErrEventNotFound {
				// ok event wasnt one of the ones asked to be parsed
			}
		}

		switch payload.(type) {
		case github.ReleasePayload:
			release := payload.(github.ReleasePayload)
			fmt.Printf("%+v", release)
		case github.PullRequestPayload:
			pullRequest := payload.(github.PullRequestPayload)
			fmt.Printf("%+v", pullRequest)
		}
	})

	http.ListenAndServe(":3000", nil)
}
