package arguments

import (
    "errors"
    "net/url"
)

func ParseArguments(args ...string) (*Arguments, error) {
    if len(args) == 0 {
        return nil, errors.New("url not provided")
    }

    arguments := NewArguments()

    rawUrl := args[0]
    parsedUrl, err := url.Parse(rawUrl)
    if err != nil {
        return nil, err
    }
    arguments.URL = parsedUrl

    return arguments, nil
}
