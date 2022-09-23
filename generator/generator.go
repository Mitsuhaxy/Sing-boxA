package generator

import (
	"io/ioutil"
	"os"
)

func Generator() (isSuccess bool) {
	if Create() {
		return false
	}
	content := []byte(`Go is an open source programming language that makes is easy to build simple,reliable,and efficient software`)
	err := ioutil.WriteFile("a.txt", content, 0777)
	return err == nil

}

func Create() (isSuccess bool) {
	file, err := os.Create("sing-box.json")
	if err == nil {
		return true
	}
	defer file.Close()
	return false
}
