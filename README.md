# slasher
--
    import "github.com/autom8ter/slasher"


## Usage

#### type HandlerFunc

```go
type HandlerFunc func(s *Slasher, client *slack.Client, command *slack.SlashCommand) (interface{}, error)
```

A handler func is run against an incoming slack slash command. It is up to the
user of this library to define their own handlers

#### type Slasher

```go
type Slasher struct {
}
```

Slasher holds a slack client, a map of functions map[string]HandlerFunc and an
array of allowed users

#### func  NewSlasher

```go
func NewSlasher(token string, allowedUsers []string) *Slasher
```
Creates a newe slasher instance

#### func (*Slasher) AddHandler

```go
func (s *Slasher) AddHandler(command string, function HandlerFunc)
```
Adds a slack slash command handler

#### func (*Slasher) BashScipt

```go
func (s *Slasher) BashScipt(script string) ([]byte, error)
```
runs exec.Command("/bin/sh", "-c", script)

#### func (*Slasher) DeleteHandler

```go
func (s *Slasher) DeleteHandler(command string)
```
Deletes a handler delete(s.functions, command)

#### func (*Slasher) Error

```go
func (s *Slasher) Error(w http.ResponseWriter, err error)
```
Writes a wrapped Slasher error to the response

#### func (*Slasher) Exec

```go
func (s *Slasher) Exec(args ...string) ([]byte, error)
```
runs exec.Command(args[0], args[1:]...)

#### func (*Slasher) Exists

```go
func (s *Slasher) Exists(command string) bool
```
Checks if a function exists for the command

#### func (*Slasher) HandlerFunc

```go
func (s *Slasher) HandlerFunc() http.HandlerFunc
```

#### func (*Slasher) JSON

```go
func (s *Slasher) JSON(w http.ResponseWriter, obj interface{})
```
Writes pretty json to the response

#### func (*Slasher) Python3Script

```go
func (s *Slasher) Python3Script(cmd string) ([]byte, error)
```
runs exec.Command("python3", "-c", cmd)

#### func (*Slasher) ShellScript

```go
func (s *Slasher) ShellScript(script string) ([]byte, error)
```
exec.Command("/bin/sh", "-c", script)
