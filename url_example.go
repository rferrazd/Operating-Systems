package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func download(url string) {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
	bodyBytes, err := io.ReadAll(resp.Body) //reading in the bytes
	if err != nil {
		fmt.Println(err)
	}
	bodyString := string(bodyBytes)
	// TO DO: change print statement to write to a file
	// ----->   fmt.Println(bodyString)  CHANGE THIS
	f1, err := os.Create("data.html") // return tuple with whatt you actually want from the function
	if err != nil {
		// print error
		fmt.Println(err)
	}
	// use defer because if program were to take on forever file handlers would be open. close connection to file so that operating system can reclaim file handlers
	defer f1.Close()
	f1.Write(bodyBytes)
	f1.Sync()

	// write the strings in another file
	f2, err := os.Create("string.html")
	if err != nil {
		fmt.Println(err)
	}
	defer f2.Close()
	f2.WriteString(bodyString)
	f2.Sync()
}

func main() {
	url := os.Args[1]
	download(url)
}
