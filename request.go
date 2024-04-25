package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Secret struct {
	SecretName string `json:"secret_name"`
	Key        string `json:"key"`
	Value      string `json:"value"`
}

func Request(domain, token, secretName string) ([]Secret, error) {
	var secrets []Secret

	if domain[len(domain)-1] == '/' {
		domain = domain[:len(domain)-1]
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/secret/%s", domain, secretName), nil)
	if err != nil {
		return secrets, err
	}

	req.Header.Set("Authorization", token)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return secrets, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return secrets, err
	}

	if err := json.Unmarshal(body, &secrets); err != nil {
		return secrets, err
	}

	return secrets, nil
}
