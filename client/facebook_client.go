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

	user := &domain.User{
		FacebookId: res.Get("id").(string),
		Name:       res.Get("name").(string),
	}

	if email := res.Get("email"); email != nil {
		user.Email = email.(string)
	}

	if picture := res.Get("picture.data.url"); picture != nil {
		user.Picture = picture.(string)
	}

	return user, nil
}
