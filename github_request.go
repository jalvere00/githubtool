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

type PullRequest struct {
	Title  string `json:"title"`
	Number int    `json:"number"`
	State  string `json:"state"`
}

var token = flag.String("token", "", "Using for Authentication of private repos.")
var apiVersion = flag.String("api-version", "2022-11-28", "GitHub version your communicating with.")

func GetRepoRelease(user, repository string) ([]Release, error) {
	var releases []Release

	response, err := makeGetRequest(user, repository, "releases")
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

func GetRepoPull(user, repository string) ([]PullRequest, error) {
	var pull []PullRequest

	response, err := makeGetRequest(user, repository, "pulls")
	if err != nil {
		fmt.Println("Error making response: ", err)
		return nil, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&pull)
	if err != nil {
		fmt.Println("Error decoding JSON response: ", err)
		return nil, err
	}

	return pull, nil
}

func createGetRequest(user, repository, api string) (*http.Request, error) {
	url := base + fmt.Sprintf("%s/%s/%s", user, repository, api)

	request, err := http.NewRequest("GET", url, nil)

	if err == nil && *token != "" {
		request.Header.Set("Authorization", *token)
	}
	request.Header.Set("X-GitHub-Api-Version", *apiVersion)

	return request, err
}

func makeGetRequest(user, repository, api string) (*http.Response, error) {
	client := &http.Client{}

	request, err := createGetRequest(user, repository, api)
	if err != nil {
		fmt.Println("Error creating request: ", err)
		return nil, err
	}

	// ToDo(jaalvere00): check status code.
	response, err := client.Do(request)
	return response, err
}
