package main

import (
	"Sing-boxA/api"
	"fmt"
	"os"
)

const SBA_VERSION string = "b1.0"

func main() {
	switch os.Args[1] {
	case "run":
		{
			api.RunApi()
		}
	case "version":
		fmt.Println(SBA_VERSION)
	case "help":
		fmt.Println("YOU DONT NEED HELP!")
	}
}
