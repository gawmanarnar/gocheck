package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type results struct {
	XMLName xml.Name `xml:"results"`
	Errder  errors   `xml:"errors"`
}

type errors struct {
	XMLName xml.Name `xml:"errors"`
	Errs    []err    `xml:"error"`
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

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		return
	}

	v := results{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < len(v.Errder.Errs); i++ {
		fmt.Printf("%s:%d - %s\n", v.Errder.Errs[i].Loc.File, v.Errder.Errs[i].Loc.Line, v.Errder.Errs[i].Msg)
	}
}
