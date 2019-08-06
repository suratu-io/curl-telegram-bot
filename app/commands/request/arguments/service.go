package arguments

import (
	"errors"
	"net/url"
	"strings"
)

func stringAppendHttps(url string) string {
	newUrl := "https://" + url
	return newUrl
}

func ParseArguments(args ...string) (*Arguments, error) {
	if len(args) == 0 {
		return nil, errors.New("url not provided")
	}

	arguments := NewArguments()
	rawUrl := args[0]
	if strings.Contains(rawUrl, "https://") == false {
		rawUrl = stringAppendHttps(rawUrl)
	}
	parsedUrl, err := url.Parse(rawUrl)
	if err != nil {
		return nil, err
	}
	arguments.URL = parsedUrl

	return arguments, nil
}
