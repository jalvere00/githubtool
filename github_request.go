package main

import (
	"net/http"
	"flag"
	"encoding/json"
)

const (
	base = "https://api.github.com/repos/"
)

type release struct {
	Name string
	Date string
}

var token = flag.String("token", "", "Using for Authentication of private repos.")

func GetRepoRelease(user, repository string) (release, error) {
	client := &http.Client{}
	url := base + fmt.Sprintf("%s%s/releases", user, repository)

	request, err := http.NewRequest("GET", url)
	if err := nil {
		fmt.Println("Error creating request: ", err)
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error making response: ", err)
		return nil, err
	}

	defer response.Body.Close()
	
}
