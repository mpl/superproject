package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type Camera struct {
	ID   int
	Name string
}

// TODO: make it a real test file/suite

func main() {
	cams := []Camera{
		{
			ID:   1,
			Name: "foo",
		},
		{
			ID:   2,
			Name: "bar",
		},
	}

	body, err := json.Marshal(cams)
	if err != nil {
		log.Fatal(err)
	}
	buf := bytes.NewReader(body)

	resp, err := http.Post("http://localhost:8080/cameras", "application/json", buf)
	if err != nil {
		log.Fatal(err)
	}

	println(resp.Status)
}
