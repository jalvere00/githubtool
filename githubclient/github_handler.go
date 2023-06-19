package githubclient

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

const (
	ReleaseCMD    = "releases"
	PullCMD       = "pulls"
	gitHubBaseAPI = "https://api.github.com/repos"
)

type Release struct {
	Name    string `json:"name"`
	TagName string `json:"tag_name"`
	Date    string `json:"created_at"`
}

// Todo(jaalvere00): Rename struct
type PullRequest struct {
	Title  string `json:"title"`
	Number int    `json:"number"`
	State  string `json:"state"`
}

var maxResponse = flag.Int("max", 3, "The maximum number of reponses for any request.")
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

func createGetRequest(user, repository, api string) (*http.Request, error) {
	url, err := createAPIUrl(user, repository, api)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("GET", *url, nil)
	if err != nil {
		return nil, err
	}

	setHeaderToken(request)

	return request, err
}

func createAPIUrl(user, repository, api string) (*string, error) {
	endpoint := fmt.Sprintf("/%s/%s/%s", user, repository, api)
	params := url.Values{}
	params.Add("per_page", strconv.Itoa(*maxResponse))

	urlBuilder, err := url.Parse(gitHubBaseAPI + endpoint)
	if err != nil {
		fmt.Println("There was a error creating a URL for this request.")
		return nil, err
	}

	urlBuilder.RawQuery = params.Encode()
	urlString := urlBuilder.String()

	return &urlString, nil
}

func setHeaderToken(request *http.Request) {
	if request != nil && *token != "" {
		request.Header.Set("Authorization", *token)
	}
	request.Header.Set("X-GitHub-Api-Version", *apiVersion)
}
