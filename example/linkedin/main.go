package main

import (
	"context"
	"fmt"
	"time"

	"github.com/wuriyanto48/go-social"
	"github.com/wuriyanto48/go-social/internal/linkedin"
)

//https://www.linkedin.com/oauth/v2/authorization?redirect_uri=http://localhost:8080/callback&response_type=code&client_id={client_id}&state=xwyz
func main() {
	l, err := social.New(social.Linkedin, "client_id", "client_secret", "http://localhost:8080/callback")

	if err != nil {
		fmt.Println(err)
	}

	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, 3000*time.Millisecond)

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
