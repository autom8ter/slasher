package examples

import (
	"fmt"
	"github.com/autom8ter/slasher"
	"github.com/nlopes/slack"
	"github.com/pkg/errors"
	"net/http"
	"os"
)


func helloWorld() slasher.HandlerFunc{
	return func(s *slasher.Slasher, client *slack.Client, command *slack.SlashCommand) (i interface{}, err error) {
		script := `echo "hello world!"`
		output, err := s.ShellScript(script)
		if err != nil {
			return nil,  errors.Wrapf(err, "failed to run script: %s", script)
		}
		return &slack.Message{
			Msg: slack.Msg{
				Text: string(output),
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