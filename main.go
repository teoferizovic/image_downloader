package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"image_downloader/model"
	"image_downloader/processor"
)

var conf model.Config

func init(){

	if _, err := toml.DecodeFile("./config.toml", &conf); err != nil {
		fmt.Println(err)
	}

}

func main() {

	err,imagePaths := processor.ReadFile(conf)

	if err!=nil{
		panic(err)
	}

	if len(imagePaths)!=0 {
		err = processor.DownloadImages(conf,imagePaths)
		if err != nil {
			panic(err)
		}
	}

}

