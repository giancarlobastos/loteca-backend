package client

import (
	fb "github.com/huandu/facebook/v2"
)

type FacebookClient struct {
}

func NewFacebookClient() *FacebookClient {
	return &FacebookClient{}
}

func (c *FacebookClient) ValidateUser(token string) (string, error) {
	res, _ := fb.Get("/me", fb.Params{
		"access_token": token,
	})
	return res["id"].(string), nil
}
