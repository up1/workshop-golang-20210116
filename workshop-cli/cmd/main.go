package main

import (
	"cli"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Initial dependencies
	store := cli.FileStore{Filename: "demo.json"}
	userService := cli.UserService{Store: &store}

	// Parser input from command line
	command := os.Args[1]
	switch command {
	case "add":
		name := os.Args[2]
		age := os.Args[3]
		iage, _ := strconv.Atoi(age)
		userService.AddNew(name, iage)
	case "list":
		res := userService.ListAll()
		fmt.Println(res)
	}
}
