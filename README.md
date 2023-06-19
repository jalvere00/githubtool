# GitHubTool

A tool for viewing the most recent pull request and release versions of a GitHub repo.

## Installing
Requires [Go][def]

### Linux
More information on setting up your go environment can be found at [link](https://www.geeksforgeeks.org/how-to-install-a-package-with-go-get/)

Run the following commands to install the CLI tool.
```bash
# Download the package
go get github.com/jalvere00/githubtool

# Compile and install package
go install -v github.com/jalvere00/GitHubTool

# Set execute path in alias
alias githubtool=~/go/bin/githubtool
```


## How To Use
Examples:
- `githubtool pulls haccer subjack` 

Usage:
```bash
githubtool <command> [username] [repository]
```

Commands:
- `pulls` fetch to most recent pulls for a repository.
- `releases` fetch to most recent release for the repository.

Options:
- `--token` a security token which allows you to access a private reposity.
- `--max` the maximum size of the reponse from github's api.
- `--api-version` the verison of github's api the tool will be communicating with.

[def]: https://go.dev/dl/
