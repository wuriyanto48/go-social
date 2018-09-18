package google

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/wuriyanto48/go-social/api"
	"github.com/wuriyanto48/go-social/internal"
)

const (
	// DefaultAuthURI default Authorization URI for Google
	DefaultAuthURI = "https://accounts.google.com/o/oauth2/auth"

	// DefaultTokenURI default Token URI for Google
	DefaultTokenURI = "https://accounts.google.com/o/oauth2/token"

	// DefaultAPIRURI default API URI for Google
	DefaultAPIRURI = "https://www.googleapis.com/oauth2/v2/userinfo"
)

// Google struct
type Google struct {
	ClientID     string
	ClientSecret string
	AuthURI      string
	APIURI       string
	TokenURI     string
	RedirectURI  string
	Token        string
	httpClient   *internal.HTTPClient
}

// New function, Google's Constructor
func New(clientID, clientSecret, redirectURI string) *Google {
	httpClient := internal.NewHTTPClient()
	return &Google{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURI:  redirectURI,
		httpClient:   httpClient,
		AuthURI:      DefaultAuthURI,
		APIURI:       DefaultAPIRURI,
		TokenURI:     DefaultTokenURI,
	}
}

// GetAuthURI function
func (g *Google) GetAuthURI() (string, error) {
	return "", nil
}

// GetAccessToken function
func (g *Google) GetAccessToken(ctx context.Context, authorizationCode string) error {

	if g.ClientID == "" {
		return internal.NewErrorEmptyValue("client id")
	}

	if g.ClientSecret == "" {
		return internal.NewErrorEmptyValue("client secret")
	}

	if g.RedirectURI == "" {
		internal.NewErrorEmptyValue("redirect uri")
	}

	form := url.Values{}
	form.Add("grant_type", "authorization_code")
	form.Add("code", authorizationCode)
	form.Add("client_id", g.ClientID)
	form.Add("client_secret", g.ClientSecret)
	form.Add("redirect_uri", g.RedirectURI)

	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	var response struct {
		AccessToken      string `json:"access_token"`
		ExpiresIn        int    `json:"expires_in"`
		RefreshToken     string `json:"refresh_token"`
		Scope            string `json:"scope"`
		TokenType        string `json:"token_type"`
		IDToken          string `json:"id_token"`
		Error            string `json:"error"`
		ErrorDescription string `json:"error_description"`
	}

	err := g.httpClient.Execute(ctx, "POST", g.TokenURI, strings.NewReader(form.Encode()), &response, headers)

	if err != nil {
		return err
	}

	// prevent error
	if len(response.Error) > 0 {
		return errors.New(response.ErrorDescription)
	}

	g.Token = response.AccessToken

	return nil
}

// GetUser function
func (g *Google) GetUser(ctx context.Context) (api.Result, error) {

	if g.Token == "" {
		return nil, internal.NewErrorEmptyValue("access token")
	}

	headers := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", g.Token),
	}

	var response User

	err := g.httpClient.Execute(ctx, "GET", g.APIURI, nil, &response, headers)

	if err != nil {
		return nil, err
	}

	if response.Error != nil {
		return nil, errors.New(response.Error.Message)
	}

	return &response, nil
}
