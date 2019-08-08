package arguments

import (
	"errors"
	"net/url"
	"strings"
)

const (
	schemaSeparator = "://"
	defaultSchema   = "https"
)

func ParseArguments(args ...string) (*Arguments, error) {
	if len(args) == 0 {
		return nil, errors.New("url not provided")
	}

	arguments := NewArguments()

	rawUrl := validateRawUrl(args[0])
	parsedUrl, err := url.Parse(rawUrl)
	if err != nil {
		return nil, err
	}
	arguments.URL = parsedUrl

	return arguments, nil
}

func validateRawUrl(url string) string {
	schemaEnd := strings.Index(url, schemaSeparator)
	if schemaEnd < 0 {
		return strings.Join([]string{defaultSchema, url}, schemaSeparator)
	}
	return url
}
