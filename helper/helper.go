package helper

import (
	"net/url"
	"strings"
)

func BuildFileName(fullUrlFile string) (string,error){

	fileUrl, err := url.Parse(fullUrlFile)

	if err != nil {
		return "",err
	}

	path := fileUrl.Path
	segments := strings.Split(path, "/")

	fileName := segments[len(segments)-1]

	return fileName, nil
}
