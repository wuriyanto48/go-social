package main

import (
	"context"
	"fmt"

	"github.com/wuriyanto48/go-social"
	"github.com/wuriyanto48/go-social/pkg/facebook"
)

//https://www.facebook.com/dialog/oauth?client_id={your_client_id}&redirect_uri=http://localhost:8080/callback&response_type=code
func main() {
	f, err := social.New(social.Facebook, "client_id", "client_secret", "", "http://localhost:8080/callback", "", 0)

	if err != nil {
		fmt.Println(err)
	}

	ctx := context.Background()

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
