package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
)

type results struct {
	XMLName xml.Name `xml:"results"`
	Errors  errors   `xml:"errors"`
}

type errors struct {
	XMLName xml.Name `xml:"errors"`
	Errors  []err    `xml:"error"`
}

type err struct {
	XMLName xml.Name `xml:"error"`
	Msg     string   `xml:"msg,attr"`
	Loc     location `xml:"location"`
}

type location struct {
	XMLName xml.Name `xml:"location"`
	File    string   `xml:"file,attr"`
	Line    int      `xml:"line,attr"`
}

func getResults(file string) errors {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return errors{}
	}

	result := results{}
	err = xml.Unmarshal(data, &result)
	if err != nil {
		fmt.Println(err)
		return errors{}
	}

	return result.Errors
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Missing file name parameter")
		return
	}

	file1 := getResults(os.Args[1])
	file2 := getResults(os.Args[2])

	fmt.Println(reflect.DeepEqual(file1, file2))
}
