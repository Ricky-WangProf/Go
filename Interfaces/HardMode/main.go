package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing parameter, please provide enough parameter.")
		return
	}

	data, err := ioutil.ReadFile(os.Args[1])

	if err != nil {
		fmt.Println("File has a problem: ", os.Args[1])
		fmt.Println("Error message: ", err)
	}

	fmt.Println("Content in the file:")
	fmt.Println("*******************************************************")
	fmt.Println(string(data))
	fmt.Println("*******************************************************")
}
