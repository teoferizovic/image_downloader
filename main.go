package main

import (
	"fmt"
	"image_downloader/helper"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	url := "https://i.ytimg.com/vi/wJyWtat3kQo/maxresdefault.jpg"

	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	fileName,err := helper.BuildFileName(url)

	if err != nil {
		log.Fatal(err)
	}

	//open a file for writing
	localUrl := "/home/teo/go/src/image_downloader/images/"+fileName
	file, err := os.Create(localUrl)
	if err != nil {
		log.Fatal(err)
	}

	// Use io.Copy to just dump the response body to the file. This supports huge files
	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Fatal(err)
	}

	file.Close()
	fmt.Println("Success!")
}

