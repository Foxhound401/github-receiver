package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/webhooks/v6/github"
)

const (
	path = "/webhooks"
)

func main() {

	hook, _ := github.New(github.Options.Secret("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"))

	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		payload, err := hook.Parse(r, github.RepositoryEvent, github.CreateEvent)

		// j, _ := json.MarshalIndent(payload, "", "\t")
		// fmt.Println(string(j))

		if err != nil {
			fmt.Printf("%+v", err)
			if err == github.ErrEventNotFound {
				// ok event wasnt one of the ones asked to be parsed
				fmt.Printf("Not found well well ")
			}
		}

		switch payload.(type) {
		case github.CreatePayload:
			createPayload := payload.(github.CreatePayload)
			j, _ := json.MarshalIndent(createPayload, "", "\t")

			fmt.Println(j)
		}
	})

	fmt.Printf("Server is running at port 3000\n")
	http.ListenAndServe(":3000", nil)
}
