package main

import (
	"fmt"
	"os"
)

const (
	ReleaseCMD = "releases"
	PullCMD    = "pulls"
)

func main() {
	argsWithProg := os.Args[1:]
	if len(argsWithProg) != 3 {
		// Todo(jaalvere00): create more information output.
		fmt.Println("Not enough arguments")
	}

	// Todo(jaalvere00): Create Error message when a command is enter that doesn't exist.
	if argsWithProg[0] == ReleaseCMD {
		releases, err := GetRepoRelease(argsWithProg[1], argsWithProg[2])

		if err != nil {
			return
		}
		fmt.Println("Here are the most recent releases")
		for _, r := range releases {
			// Todo(jaalvere00): Chhange Date to a name that more discribse what is represents.
			fmt.Printf("Name:%s\nDate:%s\n\n", r.Name, r.Date)
		}
	}

	if argsWithProg[0] == PullCMD {
		pulls, err := GetRepoPull(argsWithProg[1], argsWithProg[2])
		if err != nil {
			return
		}
		fmt.Println("Here are the most recent pull request")
		for _, p := range pulls {
			fmt.Printf("Title:%s\nNumber:%d\nState:%s\n\n", p.Title, p.Number, p.State)
		}
	}
}
