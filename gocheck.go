package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

func stripLineNumbers(file string) bool {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return false
	}

	re := regexp.MustCompile(`\Wline="(.*?)"`)
	res := re.ReplaceAll(data, []byte(""))

	err = ioutil.WriteFile(os.Args[1], res, 0644)
	if err != nil {
		return false
	}

	return true
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Missing file name parameter")
		return
	}

	stripLineNumbers(os.Args[1])
	stripLineNumbers(os.Args[2])
}
