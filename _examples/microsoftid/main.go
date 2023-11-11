package main

import (
	"context"
	"fmt"

	"github.com/wuriyanto48/go-social"
	"github.com/wuriyanto48/go-social/pkg/microsoftid"
)

//https://login.microsoftonline.com/{your_tenant_id}/oauth2/v2.0/authorize?client_id={your_client_id}&response_type=code&scope=https://graph.microsoft.com/User.Read&redirect_uri=http://localhost:3000/callback
func main() {
	g, err := social.New(social.MicrosoftID, "client_id", "client_secret", "tenant_id", "http://localhost:3000/callback", "https://graph.microsoft.com/User.Read", 0)

	if err != nil {
		fmt.Println(err)
	}

	ctx := context.Background()

	fmt.Println(g.GetAuthURI())

	err = g.GetAccessToken(ctx, "code")

	if err != nil {
		fmt.Println(err)
		return
	}

	result, err := g.GetUser(ctx)

	if err != nil {
		fmt.Println(err)
		return
	}

	user, _ := result.(*microsoftid.User)

	fmt.Println(user)
}
