package utils

import (
    "io"
    "net/http"
)

func NewCommonRequest(method, url string, body io.Reader) (*http.Request, error) {
    request, err := http.NewRequest(method, url, body)
    if err != nil {
        return request, err
    }

    request.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36")
    request.Header.Set("Content-Type", "application/json")

    return request, err
}
