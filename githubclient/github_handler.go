package githubclient

import (
	"encoding/json"
	"errors"
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

type GitAPIHandler struct {
	Client  *http.Client
	BaseUrl string
}

type Release struct {
	Name    string `json:"name"`
	TagName string `json:"tag_name"`
	Date    string `json:"created_at"`
}

type Pull struct {
	Title  string `json:"title"`
	Number int    `json:"number"`
	State  string `json:"state"`
}

var maxResponse = flag.Int("max", 3, "The maximum number of reponses for any request.")
var token = flag.String("token", "", "Using for Authentication of private repos.")
var apiVersion = flag.String("api-version", "2022-11-28", "GitHub version your communicating with.")

func (handler GitAPIHandler) GetRepoRelease(user, repository string) ([]Release, error) {
	var releases []Release

	response, err := handler.makeGetRequest(user, repository, "releases")
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&releases)
	if err != nil {
		return nil, err
	}

	return releases, nil
}

func (handler GitAPIHandler) GetRepoPull(user, repository string) ([]Pull, error) {
	var pull []Pull

	response, err := handler.makeGetRequest(user, repository, "pulls")
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&pull)
	if err != nil {
		return nil, err
	}

	return pull, nil
}

func (handler GitAPIHandler) makeGetRequest(user, repository, api string) (*http.Response, error) {
	request, err := handler.createGetRequest(user, repository, api)
	if err != nil {
		return nil, err
	}
	response, err := handler.Client.Do(request)

	if valid, err := checkResponse(response); !valid {
		return nil, err
	}
	return response, err
}

func (handler GitAPIHandler) createGetRequest(user, repository, api string) (*http.Request, error) {
	url, err := handler.createAPIUrl(user, repository, api)

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

func (handler GitAPIHandler) createAPIUrl(user, repository, api string) (*string, error) {
	endpoint := fmt.Sprintf("/%s/%s/%s", user, repository, api)
	params := url.Values{}
	params.Add("per_page", strconv.Itoa(*maxResponse))

	urlBuilder, err := url.Parse(handler.BaseUrl + endpoint)
	if err != nil {
		fmt.Println("There was a error creating a URL for this request.")
		return nil, err
	}

	urlBuilder.RawQuery = params.Encode()
	urlString := urlBuilder.String()

	return &urlString, nil
}

func CreatedGitAPIHandler() GitAPIHandler {
	return GitAPIHandler{&http.Client{}, gitHubBaseAPI}
}

func checkResponse(response *http.Response) (bool, error) {
	if response.StatusCode == 404 {
		return false, errors.New("resource not found, check your username an repository name and try again")
	} else if response.StatusCode != 200 {
		return false, errors.New("unknown error occured while making request to githubs API")
	}
	return true, nil
}

func setHeaderToken(request *http.Request) {
	if request != nil && *token != "" {
		request.Header.Set("Authorization", *token)
	}
	request.Header.Set("X-GitHub-Api-Version", *apiVersion)
}
