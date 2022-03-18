package client

import (
	"log"

	"github.com/giancarlobastos/loteca-backend/domain"
	fb "github.com/huandu/facebook/v2"
)

type FacebookClient struct {
}

func NewFacebookClient() *FacebookClient {
	return &FacebookClient{}
}

func (c *FacebookClient) GetUser(token string) (*domain.User, error) {
	res, err := fb.Get("/me", fb.Params{
		"access_token": token,
		"fields":       "id,name,picture,email",
	})

	if err != nil {
		log.Printf("Error [GetUser]: %v - [%v]", err, token)
		return nil, err
	}

	return &domain.User{
		FacebookId: res.Get("id").(string),
		Name:       res.Get("name").(string),
		Email:      res.Get("email").(string),
		Picture:    res.Get("picture.data.url").(string),
	}, nil
}
