package parser

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

func parse(url string) string {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := html.Parse(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(body.Data)

	return "body"
}
