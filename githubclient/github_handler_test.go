package githubclient

import (
	"flag"
	"fmt"
	"testing"
)

func TestCreateGetRequest(t *testing.T) {
	username := "username"
	repo := "repository"
	path := "path"

	flag.Set("token", "")

	request, err := createGetRequest(username, repo, path)
	if err != nil {
		t.Errorf("Error when creating request: %s\n", err)
	}
	if request.Header.Get("Authorization") != "" {
		t.Errorf("Error Authorization token should not be set")
	}
}

func TestSetHeaderToken(t *testing.T) {
	username := "username"
	repo := "repository"
	path := "path"
	token := "abc123"

	flag.Set("token", token)

	request, _ := createGetRequest(username, repo, path)

	if request.Header.Get("Authorization") != token {
		t.Errorf("Error in request Authorization, was Expecting: %s, Actual found: %s", token, request.Header.Get("Authorization"))
	}
}

func TestCreateAPIUrl(t *testing.T) {
	username := "username"
	repo := "repository"
	path := "path"
	expectedURL := fmt.Sprintf("%s/%s/%s/%s?per_page=%d", gitHubBaseAPI, username, repo, path, 3)
	receivedUrl, err := createAPIUrl(username, repo, path)
	if err != nil {
		t.Errorf("Error when creating API URL: %s\n", err)
	}
	if expectedURL != *receivedUrl {
		t.Errorf("Failed to create URL. Was Expecting %s, Created %s", expectedURL, *receivedUrl)
	}
}
