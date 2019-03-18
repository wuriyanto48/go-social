package github

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/wuriyanto48/go-social/api"
	"github.com/wuriyanto48/go-social/pkg"
)

const (
	// DefaultAuthURI default Authorization URI for Github
	DefaultAuthURI = "https://github.com/login/oauth/authorize"

	// DefaultTokenURI default Token URI for Github
	DefaultTokenURI = "https://github.com/login/oauth/access_token"

	// DefaultAPIRURI default API URI for Github
	DefaultAPIRURI = "https://api.github.com/user"
)

// Github struct
type Github struct {
	ClientID     string
	ClientSecret string
	AuthURI      string
	APIURI       string
	TokenURI     string
	RedirectURI  string
	Token        string
	httpClient   *pkg.HTTPClient
}

// New function, Github's Constructor
func New(clientID, clientSecret, redirectURI string) *Github {
	httpClient := pkg.NewHTTPClient()
	return &Github{
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
func (g *Github) GetAuthURI() (string, error) {
	return "", nil
}

// GetAccessToken function
func (g *Github) GetAccessToken(ctx context.Context, authorizationCode string) error {

	if g.ClientID == "" {
		return pkg.NewErrorEmptyValue("client id")
	}

	if g.ClientSecret == "" {
		return pkg.NewErrorEmptyValue("client secret")
	}

	if g.RedirectURI == "" {
		pkg.NewErrorEmptyValue("redirect uri")
	}

	form := url.Values{}
	form.Add("grant_type", "authorization_code")
	form.Add("code", authorizationCode)
	form.Add("client_id", g.ClientID)
	form.Add("client_secret", g.ClientSecret)
	form.Add("redirect_uri", g.RedirectURI)

	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
		"Accept":       "application/json",
	}

	var response struct {
		AccessToken      string `json:"access_token"`
		TokenType        string `json:"token_type"`
		Scope            string `json:"scope"`
		Error            string `json:"error"`
		ErrorDescription string `json:"error_description"`
		ErrorURI         string `json:"error_uri"`
	}

	err := g.httpClient.Execute(ctx, "POST", g.TokenURI, strings.NewReader(form.Encode()), &response, headers)

	if err != nil {
		return err
	}

	if len(response.Error) > 0 {
		return errors.New(response.ErrorDescription)
	}

	g.Token = response.AccessToken

	return nil
}

// GetUser function
func (g *Github) GetUser(ctx context.Context) (api.Result, error) {

	if g.Token == "" {
		return nil, pkg.NewErrorEmptyValue("access token")
	}

	headers := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", g.Token),
	}

	var response User

	err := g.httpClient.Execute(ctx, "GET", g.APIURI, nil, &response, headers)

	if err != nil {
		return nil, err
	}

	if len(response.Message) > 0 {
		return nil, errors.New(response.Message)
	}

	return &response, nil
}
