package main

import (
	"fmt"
	"strings"

	"github.com/sethvargo/go-githubactions"
)

func main() {
	secret := githubactions.GetInput("secret")
	if secret == "" {
		githubactions.Fatalf("missing input 'secret'")
	}
	// githubactions.AddMask(secret)

	// format: http(s)://0.0.0.0:8080@token or http(s)://hostname(:port)@token
	// things between () are optional
	apiUrl := githubactions.GetInput("api_url")
	if apiUrl == "" {
		githubactions.Fatalf("missing input 'api_url'")
	}

	domain := strings.Split(apiUrl, "@")[0]
	if len(domain) != 2 {
		githubactions.Fatalf("invalid api_url format")
	}

	token := strings.Split(apiUrl, "@")[1]
	if token == "" {
		githubactions.Fatalf("missing token in api_url")
	}

	// res, err := Request("http://localhost:8080", "123", "secretName")
	res, err := Request(domain, token, secret)
	if err != nil {
		githubactions.Fatalf("error: %v", err)
	}

	for _, secret := range res {
		fmt.Printf("%+v\n", secret)
		githubactions.SetEnv(secret.Key, secret.Value)
		githubactions.AddMask(secret.Value)
	}

	githubactions.SetOutput("output", "done")
}
