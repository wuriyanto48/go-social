package main

import (
	"context"
	"fmt"

	"github.com/wuriyanto48/go-social"
	"github.com/wuriyanto48/go-social/pkg/github"
)

//https://github.com/login/oauth/authorize?client_id={your_client_id}&redirect_uri=http://localhost:8080/callback&scope=user,repo
func main() {
	g, err := social.New(social.Github, "client_id", "client_secret", "", "http://localhost:8080/callback", "", 0)

	if err != nil {
		fmt.Println(err)
	}

	ctx := context.Background()

	err = g.GetAccessToken(ctx, "code")

	if err != nil {
		fmt.Println(err)
	}

	result, err := g.GetUser(ctx)

	if err != nil {
		fmt.Println(err)
	}

	user, _ := result.(*github.User)

	fmt.Println(user.Bio)
}
