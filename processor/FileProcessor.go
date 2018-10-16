package processor

import (
	"encoding/csv"
	"fmt"
	"image_downloader/helper"
	"image_downloader/model"
	"io"
	"net/http"
	"os"
)

type CsvLine struct {
	ImageUrl string
}

func ReadFile(conf model.Config)  (error,[]string) {

	var imagePaths []string

	filename := conf.FilePath+"images.csv"

	// Open CSV file
	f, err := os.Open(filename)
	if err != nil {
		return err, nil
	}
	defer f.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return err, nil
	}

	// Loop through lines & turn into object
	for _, line := range lines {
		data := CsvLine{
			ImageUrl: line[0],
		}
		 imagePaths = append(imagePaths,data.ImageUrl)
	}

	//remove first elemnt from array/slice
	imagePaths = imagePaths[1:]

	//fmt.Println(imagePaths)
	return nil,imagePaths
}

func DownloadImages(conf model.Config,imagePaths[]string) error {

	for _, url := range imagePaths {

		//url := "https://i.ytimg.com/vi/wJyWtat3kQo/maxresdefault.jpg"
		response, err := http.Get(url)

		if err != nil {
			return err
		}

		defer response.Body.Close()

		fileName, err := helper.BuildFileName(url)

		if err != nil {
			return err
		}

		//open a file for writing
		localUrl := conf.ImagePath + fileName

		file, err := os.Create(localUrl)
		if err != nil {
			return err
		}

		// Use io.Copy to just dump the response body to the file. This supports huge files
		_, err = io.Copy(file, response.Body)
		if err != nil {
			return err
		}

		file.Close()
	}

	fmt.Println("Success!")
	return nil

}
