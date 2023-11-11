package main

import (
	"context"
	"fmt"

	"github.com/wuriyanto48/go-social"
	"github.com/wuriyanto48/go-social/pkg/linkedin"
)

//https://www.linkedin.com/oauth/v2/authorization?redirect_uri=http://localhost:8080/callback&response_type=code&client_id={client_id}&state=xwyz
func main() {
	l, err := social.New(social.Linkedin, "client_id", "client_secret", "", "http://localhost:8080/callback", "", 0)

	if err != nil {
		fmt.Println(err)
	}

	ctx := context.Background()

	err = l.GetAccessToken(ctx, "code")

	if err != nil {
		fmt.Println(err)
	}

	result, err := l.GetUser(ctx)

	if err != nil {
		fmt.Println(err)
	}

	user, _ := result.(*linkedin.User)

	fmt.Println(user)
}
