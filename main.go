package main

import (
	"Sing-boxA/api"
	"os"
	"strconv"
	"fmt"
)

const SBA_VERSION string = "b1.0"

func main() {
	for _, args := range os.Args {
		switch args[1] {
		case "run":
			api.RunApi()
		case "version":
			fmt.Println(SBA_VERSION)
		}
	}
	
}
