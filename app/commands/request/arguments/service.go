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

// ParseArguments parses command message arguments to Arguments struct.
func ParseArguments(args ...string) (*Arguments, error) {
    if len(args) == 0 {
        return nil, errors.New("url not provided")
    }

    arguments := NewArguments()

    rawURL := validateRawURL(args[0])
    parsedURL, err := url.Parse(rawURL)
    if err != nil {
        return nil, err
    }
    arguments.URL = parsedURL

    return arguments, nil
}

func validateRawURL(url string) string {
    schemaEnd := strings.Index(url, schemaSeparator)
    if schemaEnd < 0 {
        return strings.Join([]string{defaultSchema, url}, schemaSeparator)
    }
    return url
}
