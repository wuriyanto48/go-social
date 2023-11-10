package main

import (
	"context"
	"fmt"
	"time"

	"github.com/wuriyanto48/go-social"
	"github.com/wuriyanto48/go-social/pkg/facebook"
)

//https://www.facebook.com/dialog/oauth?client_id={your_client_id}&redirect_uri=http://localhost:8080/callback&response_type=code
func main() {
	f, err := social.New(social.Facebook, "client_id", "client_secret", "", "http://localhost:8080/callback", "")

	if err != nil {
		fmt.Println(err)
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2000*time.Millisecond)

	defer func() { cancel() }()

	err = f.GetAccessToken(ctx, "code")

	if err != nil {
		fmt.Println(err)
	}

	result, err := f.GetUser(ctx)

	if err != nil {
		fmt.Println(err)
	}

	user, _ := result.(*facebook.User)

	fmt.Println(user.Picture)
}
