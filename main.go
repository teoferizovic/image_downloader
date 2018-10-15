package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"image_downloader/helper"
	"image_downloader/model"
	"io"
	"log"
	"net/http"
	"os"
)

var conf model.Config

func init(){

	if _, err := toml.DecodeFile("./config.toml", &conf); err != nil {
		fmt.Println(err)
	}
	//fmt.Printf("%#v\n", conf)
}

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
	localUrl := conf.FilePath+fileName

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

