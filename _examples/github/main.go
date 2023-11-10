package main

import (
	"context"
	"fmt"
	"time"

	"github.com/wuriyanto48/go-social"
	"github.com/wuriyanto48/go-social/pkg/github"
)

//https://github.com/login/oauth/authorize?client_id={your_client_id}&redirect_uri=http://localhost:8080/callback&scope=user,repo
func main() {
	g, err := social.New(social.Github, "client_id", "client_secret", "", "http://localhost:8080/callback", "")

	if err != nil {
		fmt.Println(err)
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2000*time.Millisecond)

	defer func() { cancel() }()

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
