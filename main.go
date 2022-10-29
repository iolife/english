package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
)

func main() {
	type Sentents struct {
		English string `json:"english,omitempty"`
		Chinese string `json:"chinese,omitempty"`
		Comment string `json:"comment,omitempty"`
	}

	file, err := os.Open("abc.json")
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(file)
	if err != nil {
		file.Close()
		panic(err)
	}
	file.Close()
	var ss []Sentents
	err = json.Unmarshal(body, &ss)
	if err != nil {
		panic(err)
	}
	for {
		fmt.Println("1. input english \n2. input chinese")
		input := ""
		fmt.Scan(&input)
		fmt.Printf("you input %s\n", input)
		if input == "1" {
			bf := bufio.NewReader(os.Stdin)
			for {
				line, _, _ := bf.ReadLine()
				if string(line) == "q" {
					break
				}
				ss = append(ss, Sentents{English: string(line)})
				continue
			}

		}
		// var chinese, english, comment string
		var chinese string
		fmt.Println("plesese input")
		fmt.Scanln(&chinese)
		if chinese == "q" {
			b, err := json.Marshal(&ss)
			if err != nil {
				panic(err)
			}
			ioutil.WriteFile("abc.json", b, fs.ModePerm)
			break
		}
	}
}
