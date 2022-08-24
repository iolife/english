package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	m:=map[string]string{
		"":""
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("input start begin game or input q over game")
	for {
	
		result, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("read error:", err)
			continue
		}
		fmt.Println("result:", result)
		if result == "q\n" || result == "quit" {
			break
		}
	}
}
