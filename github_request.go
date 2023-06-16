package main

import (
	"net/http"
	"flag"
)

const (
	base
)

var token = flag.String("token", "", "Using for Authentication of private repos.")

func GetRepoRelease(user, repository string)  {
	client := &http.Client{}

	request, err := http.NewRequest("GET", "")
}