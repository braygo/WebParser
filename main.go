package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	//This request uses default headers
	//Works just fine now but may need to make custom http client in order to change headers in the future

	//Make HTTP GET request
	response, err := http.Get("https://fluid.cx/")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	//Copy data from the response to standard output
	dataInBytes, err := ioutil.ReadAll(response.Body)
	pageContent := string(dataInBytes)

	//Finding Substring
	//loader.css

	titleStartIndex := strings.Index(pageContent, "<title>")
	if titleStartIndex == -1 {
		fmt.Println("No such element found")
		os.Exit(0)
	}

	//Offset titleStartIndex by number of characters in substring (loader)
	titleStartIndex += 7

	//Finding end index .css
	titleEndIndex := strings.Index(pageContent, "</title>")
	if titleStartIndex == -1 {
		fmt.Println("No such end element found")
		os.Exit(0)
	}

	pageString := []byte(pageContent[titleStartIndex:titleEndIndex])

	fmt.Printf("String found: %s \n", pageString)

}
