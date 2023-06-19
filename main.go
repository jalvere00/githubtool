package main

import (
	"fmt"
	"os"

	"github.com/jalvere00/githubtool/githubclient"
)

const (
	ReleaseCMD = "releases"
	PullCMD    = "pulls"
)

func main() {
	argsWithProg := os.Args[1:]

	if len(argsWithProg) != 3 {
		if len(argsWithProg) < 3 {
			fmt.Fprintln(os.Stderr, "To Few arguements entered.")
		} else if len(argsWithProg) > 3 {
			fmt.Fprintln(os.Stderr, "To Many arguements entered.")
		}
		fmt.Fprintln(os.Stderr, "githubtool <command> [username] [repository]")
		os.Exit(1)
	}

	if argsWithProg[0] == ReleaseCMD {
		releases, err := githubclient.GetRepoRelease(argsWithProg[1], argsWithProg[2])

		if err != nil {
			fmt.Fprintf(os.Stderr, "Command failed due to the following error: %s\n", err)
			os.Exit(1)
		}
		fmt.Println("Here are the most recent releases")
		for _, r := range releases {
			fmt.Printf("Name:%s\nCreate Time:%s\n\n", r.Name, r.Date)
		}
	} else if argsWithProg[0] == PullCMD {
		pulls, err := githubclient.GetRepoPull(argsWithProg[1], argsWithProg[2])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Command failed due to the following error: %s\n", err)
			os.Exit(1)
		}
		fmt.Println("Here are the most recent pull request")
		for _, p := range pulls {
			fmt.Printf("Title:%s\nNumber:%d\nState:%s\n\n", p.Title, p.Number, p.State)
		}
	} else {
		fmt.Fprintf(os.Stderr, "Error: %s is not a command.\n", argsWithProg[0])
		os.Exit(1)
	}
}
