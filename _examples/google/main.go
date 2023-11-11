package main

import (
	"context"
	"fmt"

	"github.com/wuriyanto48/go-social"
	"github.com/wuriyanto48/go-social/pkg/google"
)

//https://accounts.google.com/o/oauth2/auth?redirect_uri=http://localhost:8080/callback&response_type=code&client_id={your_client_id}&scope=https://www.googleapis.com/auth/analytics.readonly+https://www.googleapis.com/auth/userinfo.email&approval_prompt=force&access_type=offline
func main() {
	g, err := social.New(social.Google, "client_id", "client_secret", "", "http://localhost:8080/callback", "", 0)

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

	user, _ := result.(*google.User)

	fmt.Println(user)
}
