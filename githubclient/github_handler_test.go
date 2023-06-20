package githubclient

import (
	"flag"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetRepoRelease(t *testing.T) {
	username := "username"
	repo := "repository"
	path := "releases"

	url := fmt.Sprintf("/%s/%s/%s", username, repo, path)

	server := createTestServer(url, []byte(`[{"name":"v1.3.0", "created_at":"2018-12-15T02:08:59Z", "tag_name":"1.3.0"},
	{"name":"v1.2.0", "created_at":"2018-09-17T00:48:04Z", "tag_name":"1.2.0"},
	{"name":"v1.1.0", "created_at":"2018-06-25T23:18:23Z", "tag_name":"1.1.0"}]`))

	defer server.Close()

	fmt.Println("Server: ", server.URL)
	handler := GitAPIHandler{server.Client(), server.URL}

	wantResp := []Release{{"v1.3.0", "1.3.0", "2018-12-15T02:08:59Z"},
		{"v1.2.0", "1.2.0", "2018-09-17T00:48:04Z"},
		{"v1.1.0", "1.1.0", "2018-06-25T23:18:23Z"},
	}
	receiveResp, err := handler.GetRepoRelease(username, repo)
	if err != nil {
		t.Errorf("Error: %s\n", err)
	}

	for idx, received := range receiveResp {
		if received != wantResp[idx] {
			t.Errorf("Was expecting %s, but received %s", wantResp[idx], received)
		}
	}
}

func TestGetRepoPulls(t *testing.T) {
	username := "username"
	repo := "repository"
	path := "pulls"

	url := fmt.Sprintf("/%s/%s/%s", username, repo, path)

	server := createTestServer(url, []byte(`[{"title":"Title1", "number":99, "state":"open"},
	{"title":"Title2", "number":80, "state":"open"},
	{"title":"Title3", "number":33, "state":"open"}]`))

	defer server.Close()

	fmt.Println("Server: ", server.URL)
	handler := GitAPIHandler{server.Client(), server.URL}

	wantResp := []Pull{{"Title1", 99, "open"},
		{"Title2", 80, "open"},
		{"Title3", 33, "open"},
	}
	receiveResp, err := handler.GetRepoPull(username, repo)
	if err != nil {
		t.Errorf("Error: %s\n", err)
	}

	for idx, received := range receiveResp {
		if received != wantResp[idx] {
			t.Errorf("Was expecting %v, but received %v", wantResp[idx], received)
		}
	}
}

func TestCreateGetRequest(t *testing.T) {
	handler := CreatedGitAPIHandler()
	username := "username"
	repo := "repository"
	path := "path"

	flag.Set("token", "")

	request, err := handler.createGetRequest(username, repo, path)
	if err != nil {
		t.Errorf("Error when creating request: %s\n", err)
	}
	if request.Header.Get("Authorization") != "" {
		t.Errorf("Error Authorization token should not be set")
	}
}

func TestSetHeaderToken(t *testing.T) {
	handler := CreatedGitAPIHandler()
	username := "username"
	repo := "repository"
	path := "path"
	token := "abc123"

	flag.Set("token", token)

	request, _ := handler.createGetRequest(username, repo, path)

	if request.Header.Get("Authorization") != token {
		t.Errorf("Error in request Authorization, was Expecting: %s, Actual found: %s", token, request.Header.Get("Authorization"))
	}
}

func TestCreateAPIUrl(t *testing.T) {
	handler := CreatedGitAPIHandler()
	username := "username"
	repo := "repository"
	path := "path"
	expectedURL := fmt.Sprintf("%s/%s/%s/%s?per_page=%d", gitHubBaseAPI, username, repo, path, 3)
	receivedUrl, err := handler.createAPIUrl(username, repo, path)
	if err != nil {
		t.Errorf("Error when creating API URL: %s\n", err)
	}
	if expectedURL != *receivedUrl {
		t.Errorf("Failed to create URL. Was Expecting %s, Created %s", expectedURL, *receivedUrl)
	}
}

func createTestServer(pattern string, jsonResp []byte) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(createHandlerFunc(jsonResp)))
}

func createHandlerFunc(jsonResponse []byte) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		rw.Write(jsonResponse)
		rw.WriteHeader(http.StatusOK)
	}
}
