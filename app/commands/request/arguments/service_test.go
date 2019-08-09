package arguments

import (
    "github.com/stretchr/testify/require"
    "testing"
)

func TestParseArguments(t *testing.T) {
    _, err := ParseArguments()
    require.Equal(t, err.Error(), urlNotProvidedErr, "url's not provided, but func still won't return err")

    testURL := "api.example.com"
    expectedURL := defaultSchema + schemaSeparator + testURL
    args, err := ParseArguments(testURL)
    require.Nil(t, err)
    require.Equal(t, args.URL.String(), expectedURL, "default schema should be appended to url if its schema's missing")
}
