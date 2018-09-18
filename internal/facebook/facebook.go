package facebook

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
	// DefaultAuthURI default Authorization URI for Facebook
	DefaultAuthURI = "https://www.facebook.com/dialog/oauth"

	// DefaultTokenURI default Token URI for Facebook
	DefaultTokenURI = "https://graph.facebook.com/oauth/access_token"

	// DefaultAPIRURI default API URI for Facebook
	DefaultAPIRURI = "https://graph.facebook.com"
)

// Facebook struct
type Facebook struct {
	ClientID     string
	ClientSecret string
	AuthURI      string
	APIURI       string
	TokenURI     string
	RedirectURI  string
	Token        string
	httpClient   *internal.HTTPClient
}

// New function, Facebook's Constructor
func New(clientID, clientSecret, redirectURI string) *Facebook {
	httpClient := internal.NewHTTPClient()
	return &Facebook{
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
func (f *Facebook) GetAuthURI() (string, error) {
	return "", nil
}

// GetAccessToken function
func (f *Facebook) GetAccessToken(ctx context.Context, authorizationCode string) error {

	if f.ClientID == "" {
		return internal.NewErrorEmptyValue("client id")
	}

	if f.ClientSecret == "" {
		return internal.NewErrorEmptyValue("client secret")
	}

	if f.RedirectURI == "" {
		internal.NewErrorEmptyValue("redirect uri")
	}

	form := url.Values{}
	form.Add("grant_type", "authorization_code")
	form.Add("code", authorizationCode)
	form.Add("client_id", f.ClientID)
	form.Add("client_secret", f.ClientSecret)
	form.Add("redirect_uri", f.RedirectURI)

	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	var response struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		ExpiresIn   int    `json:"expires_in"`
		Error       *struct {
			Message   string `json:"message"`
			Type      string `json:"type"`
			Code      int    `json:"code"`
			FbtraceID string `json:"fbtrace_id"`
		} `json:"error,omitempty"`
	}

	err := f.httpClient.Execute(ctx, "POST", f.TokenURI, strings.NewReader(form.Encode()), &response, headers)

	if err != nil {
		return err
	}

	// prevent error
	if response.Error != nil {
		return errors.New(response.Error.Message)
	}

	f.Token = response.AccessToken

	return nil
}

// GetUser function
func (f *Facebook) GetUser(ctx context.Context) (api.Result, error) {

	if f.Token == "" {
		return nil, internal.NewErrorEmptyValue("access token")
	}

	var response User

	uri := fmt.Sprintf("%s/me?fields=id,name,email,birthday,gender,picture{height,is_silhouette,url,width}&access_token=%s", f.APIURI, f.Token)

	err := f.httpClient.Execute(ctx, "GET", uri, nil, &response, nil)

	if err != nil {
		return nil, err
	}

	if response.Error != nil {
		return nil, errors.New("get user facebook error")
	}

	return &response, nil
}
