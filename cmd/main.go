package main

import "github.com/Murat993/cli-chat/cmd/root"

func main() {
	// go build -o ./bin/my_app cmd/main.go
	// bin/my_app create user -u murat  ИЛИ  bin/my_app create user --username murat
	root.Execute()
}
