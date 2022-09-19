package run

import (
	"Sing-boxA/generator"
	"fmt"
)

func Run(command string) (isSuccess bool) {
	switch command {
	case "start":
		return StartInstance()
	case "stop":
		return StopInstance()
	case "updategeodata":
		return UpdateGeodata()
	}
	return false
}

func StartInstance() (isSuccess bool) {
	path, _ := generator.Generator()
	fmt.Println("sing-box start!")
	return true
}

func StopInstance() (isSuccess bool) {
	fmt.Println("sing-box stop!")
	return true
}

func UpdateGeodata() (isSuccess bool) {
	fmt.Println("sing-box geodate update success!!")
	return true
}
