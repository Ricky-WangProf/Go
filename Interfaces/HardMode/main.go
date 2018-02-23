package main

import (
	"fmt"
	"io"
	"os"
)

type fileContent struct{}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing parameter, please provide enough parameter.")
		return
	}

	file, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println("File has a problem: ", os.Args[1])
		fmt.Println("Error message: ", err)
		os.Exit(1)
	}

	fc := fileContent{}

	io.Copy(fc, file)
	// ******************************************************************************
	// data, err := ioutil.ReadFile(os.Args[1])

	// if err != nil {
	// 	fmt.Println("File has a problem: ", os.Args[1])
	// 	fmt.Println("Error message: ", err)
	// }

	// fmt.Println("Content in the file:")
	// fmt.Println(string(data))
	// *****************************************************************************
}

func (fileContent) Write(b []byte) (int, error) {
	fmt.Println(string(b))

	return len(b), nil
}
