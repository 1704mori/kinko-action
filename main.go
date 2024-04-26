package main

import (
	"fmt"
	"strings"

	"github.com/sethvargo/go-githubactions"
)

func main() {
	secret := githubactions.GetInput("secret")
	if secret == "" {
		githubactions.Fatalf("missing input 'secret': secret is required")
	}

	apiUrl := githubactions.GetInput("api_url")
	fmt.Printf("api_url: %s\n", apiUrl)
	if apiUrl == "" {
		githubactions.Fatalf("missing input 'api_url': api_url is required")
	}

	domain := strings.Split(apiUrl, "@")[0]
	if domain == "" {
		githubactions.Fatalf("missing domain in api_url: domain is required")
	}

	token := strings.Split(apiUrl, "@")[1]
	if token == "" {
		githubactions.Fatalf("missing token in api_url: token is required")
	}

	res, err := Request(domain, token, secret)
	if err != nil {
		githubactions.Fatalf("error making request: %v", err)
	}

	for _, secret := range res {
		fmt.Printf("%+v\n", secret)
		githubactions.SetEnv(secret.Key, secret.Value)
		githubactions.AddMask(secret.Value)
	}

	githubactions.SetOutput("output", "done")
}
