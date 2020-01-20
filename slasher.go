//go:generate godocdown -o README.md

package slasher

import (
	"encoding/json"
	"fmt"

	"github.com/nlopes/slack"
	"github.com/thoas/go-funk"

	"io"
	"net/http"
	"os"
	"os/exec"
)

//A handler func is run against an incoming slack slash command. It is up to the user of this library to define their own handlers
type HandlerFunc func(s *Slasher, client *slack.Client, command *slack.SlashCommand) (interface{}, error)

//Slasher holds a slack client, a map of functions map[string]HandlerFunc and an array of allowed users
type Slasher struct {
	client       *slack.Client
	functions    map[string]HandlerFunc
	allowedUsers []string
}

//Creates a newe slasher instance
func NewSlasher(token string, allowedUsers []string) *Slasher {
	return &Slasher{
		client:       slack.New(token),
		functions:    make(map[string]HandlerFunc),
		allowedUsers: allowedUsers,
	}
}

//Adds a slack slash command handler
func (s *Slasher) AddHandler(command string, function HandlerFunc) {
	s.functions[command] = function
}

//Checks if a function exists for the command
func (s *Slasher) Exists(command string) bool {
	_, ok := s.functions[command]
	return ok
}

//Deletes a handler delete(s.functions, command)
func (s *Slasher) DeleteHandler(command string) {
	delete(s.functions, command)
}

//Writes a wrapped Slasher error to the response
func (s *Slasher) Error(w http.ResponseWriter, err error) {
	msg := &slack.Message{
		Msg: slack.Msg{
			Text: fmt.Sprintf(":feelsbadman: Slasher error: %s", err.Error()),
		},
	}
	s.JSON(w, msg)
	return
}

//Writes pretty json to the response
func (s *Slasher) JSON(w http.ResponseWriter, obj interface{}) {
	bits, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		s.Error(w, err)
	}
	_, err = w.Write(bits)
	if err != nil {
		s.Error(w, err)
	}
	w.WriteHeader(http.StatusOK)
	return
}

//Writes the string to the response
func (s *Slasher) String(w http.ResponseWriter, response string) {
	_, err := io.WriteString(w, response)
	if err != nil {
		s.Error(w, err)
	}
	w.WriteHeader(http.StatusOK)
	return
}

func (s *Slasher) HandlerFunc() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		slash, err := slack.SlashCommandParse(r)
		if err != nil {
			s.Error(w, err)
		}
		if !funk.ContainsString(s.allowedUsers, slash.UserName) {
			msg := &slack.Message{
				Msg: slack.Msg{
					Text: fmt.Sprintf(":feelsbadman: %s is not authorized to use slasher! :wink:", slash.UserName),
				},
			}
			s.JSON(w, msg)
		}
		for command, function := range s.functions {
			if command == slash.Command {
				obj, err := function(s, s.client, &slash)
				if err != nil {
					s.Error(w, err)
				}
				s.JSON(w, obj)
			}
		}
		http.Error(w, fmt.Sprintf(":feelsbadman: slash handler %s not found", slash.Command), http.StatusNotFound)
		return
	}
}

//runs exec.Command(args[0], args[1:]...)
func (s *Slasher) Exec(args ...string) ([]byte, error) {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Env = append(cmd.Env, os.Environ()...)
	stdoutb, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("running %v: %v", cmd.Args, err)
	}
	return stdoutb, nil
}

//runs exec.Command("/bin/sh", "-c", script)
func (s *Slasher) BashScipt(script string) ([]byte, error) {
	e := exec.Command("/bin/sh", "-c", script)
	e.Env = os.Environ()
	return e.Output()
}

//exec.Command("/bin/sh", "-c", script)
func (s *Slasher) ShellScript(script string) ([]byte, error) {
	e := exec.Command("/bin/sh", "-c", script)
	e.Env = os.Environ()
	return e.Output()
}

//runs exec.Command("python3", "-c", cmd)
func (s *Slasher) Python3Script(cmd string) ([]byte, error) {
	e := exec.Command("python3", "-c", cmd)
	e.Env = os.Environ()
	return e.Output()
}

func ExampleHandler() {

}
