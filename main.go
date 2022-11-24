package main

import (
	"Sing-boxA/api"
	"os"
	"fmt"
)

const SBA_VERSION string = "b1.0"

func main() {
	
		switch os.Args[1] {
		case "run":
			api.RunApi()
		case "version":
			fmt.Println(SBA_VERSION)
		}
	
	
}
