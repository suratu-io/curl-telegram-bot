package arguments

import (
    "net/http"
    "net/url"
)

type (
    Arguments struct {
        URL     *url.URL
        Headers http.Header
    }
)

func NewArguments() *Arguments {
    return &Arguments{
        Headers: http.Header{},
    }
}
