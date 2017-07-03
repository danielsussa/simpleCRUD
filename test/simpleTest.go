package main

import (
	"fmt"
	"github.com/go-resty/resty"
)

func requestGET(url string, query map[string]string, token string) error {
	resp, err := resty.R().SetHeader("Accept", "application/file").
		SetQueryParams(query).
		SetAuthToken(token).
		Get(url)
	if err != nil {
		return err
	}
	fmt.Println("Response:", resp.String())
	return nil
}

func main() {
	requestGET("http://127.0.0.1:8080", nil, "")
}
