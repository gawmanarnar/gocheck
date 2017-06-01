package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
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

	cmd := "svn"
	args := []string{"cat", os.Args[1]}

	out, err := exec.Command(cmd, args...).Output()
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile("temp.cpp", out, 0644)
	if err != nil {
		log.Fatal(err)
	}

	cmd = "C:\\Program Files\\cppcheck\\cppcheck.exe"
	args = []string{"--force", "-j 4", "--enable=all", "--inconclusive", "--inline-suppr", "--xml", "--xml-version=2", "--std=c++03", "--suppress=cstyleCast",
		"--suppress=noExplicitConstructor", "--suppress=missingInclude", "--suppress=unmatchedSuppression", "temp.cpp"}

	f, err := os.Create("cppcheck.xml")
	if err != nil {
		log.Fatal(err)
	}

	command := exec.Command(cmd, args...)
	command.Stderr = f
	if err := command.Start(); err != nil {
		log.Fatal(err)
	}

	if err := command.Wait(); err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadFile(f.Name())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))

	err = os.Remove("temp.cpp")
	if err != nil {
		fmt.Println(err)
		return
	}

	f.Close()
	err = os.Remove("cppcheck.xml")
	if err != nil {
		fmt.Println(err)
		return
	}

	// if len(os.Args) < 3 {
	// 	fmt.Println("Missing file name parameter")
	// 	return
	// }

	// file1 := getResults(os.Args[1])
	// file2 := getResults(os.Args[2])

	// fmt.Println(reflect.DeepEqual(file1, file2))
}
