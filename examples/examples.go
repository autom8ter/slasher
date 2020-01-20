package examples

import (
	"fmt"
	"github.com/autom8ter/slasher"
	"github.com/nlopes/slack"
	"net/http"
	"os"
)


func helloWorld() slasher.HandlerFunc{
	return func(s *slasher.Slasher, client *slack.Client, command *slack.SlashCommand) (i interface{}, err error) {
		return &slack.Message{
			Msg: slack.Msg{
				Text: "Hello World!",
			},
		}, nil
	}
}

func Example() {
	slash := slasher.NewSlasher(os.Getenv("SLACK_TOKEN"), []string{"autom8ter"})
	slash.AddHandler("/hello-world", helloWorld())
	mux := http.NewServeMux()
	mux.Handle("/slasher", slash.HandlerFunc())
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Printf("failed to start server: %s", err.Error())
		os.Exit(1)
	}
}