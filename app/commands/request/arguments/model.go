package arguments

import (
    "net/http"
    "net/url"
)

type (
    // Arguments is request configurations parsed form command message.
    Arguments struct {
        URL     *url.URL
        Headers http.Header
    }
)

// NewArguments returns new Arguments struct.
func NewArguments() *Arguments {
    return &Arguments{
        Headers: http.Header{},
    }
}
