package main

import (
	"context"
	"fmt"
	"time"

	"github.com/wuriyanto48/go-social/internal/facebook"
)

//https://www.facebook.com/dialog/oauth?client_id={your_client_id}&redirect_uri=http://localhost:8080/callback&response_type=code
func main() {
	f := facebook.New("client_id", "client_secret", "http://localhost:8080/callback")

	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, 2000*time.Millisecond)

	err := f.GetAccessToken(ctx, "AQD6WmtbIe1HylOI2xOCNeYs-3rBJWssMYSHw2Kin03Amy0OrS7mgUJbmWVkhhWNNCGfkbXLKwx15F0CISxsq7aXEEK1jjgq8XZsI_Q4di00Vl9xkC3P__cwBNMKul8WNiIqPlRCKwG5vYebrcZ9o0J3iPg2WMBX9UBa2lEnXPyOfecj8EJf4ZDWtsOJT3YT216YvdWxyOy_i8q1Gq2h_vTr537QbhbgtmXQ00ENpw6YWbA5Yg-5vhlx79nF26X4OXoL7KaniePe_0n-Nt2fHD0FQ8oPpNuivLoao7XYBp0yS-y79NxZD7vrau07NWDms_q69JUxD_UxGq31VEsA7OhHSI0_prf1YcWIOk7djFj5bg#_=_")

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
