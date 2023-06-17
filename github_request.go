package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
)

const (
	base = "https://api.github.com/repos/"
)

type Release struct {
	Name    string `json:"name"`
	TagName string `json:"tag_name"`
	Date    string `json:"created_at"`
}

var token = flag.String("token", "", "Using for Authentication of private repos.")

func GetRepoRelease(user, repository string) ([]Release, error) {
	var releases []Release

	client := &http.Client{}
	url := base + fmt.Sprintf("%s/%s/releases", user, repository)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request: ", err)
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error making response: ", err)
		return nil, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&releases)
	if err != nil {
		fmt.Println("Error decoding JSON response: ", err)
		return nil, err
	}

	return releases, nil
}
