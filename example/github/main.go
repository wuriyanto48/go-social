package main

import (
	"context"
	"fmt"
	"time"

	"github.com/wuriyanto48/go-social/internal/github"
)

//https://github.com/login/oauth/authorize?client_id={your_client_id}&redirect_uri=http://localhost:8080/callback&scope=user,repo
func main() {
	g := github.New("client_id", "client_secret", "http://localhost:8080/callback")

	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, 3000*time.Millisecond)

	err := g.GetAccessToken(ctx, "code")

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
