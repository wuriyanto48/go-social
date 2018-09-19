package linkedin

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
	// DefaultAuthURI default Authorization URI for Linkedin
	DefaultAuthURI = "https://www.linkedin.com/oauth/v2/authorization"

	// DefaultTokenURI default Token URI for Linkedin
	DefaultTokenURI = "https://www.linkedin.com/oauth/v2/accessToken"

	// DefaultAPIRURI default API URI for Linkedin
	DefaultAPIRURI = "https://api.linkedin.com/v1/people/~"
)

// Linkedin struct
type Linkedin struct {
	ClientID     string
	ClientSecret string
	AuthURI      string
	APIURI       string
	TokenURI     string
	RedirectURI  string
	Token        string
	State        string
	httpClient   *internal.HTTPClient
}

// New function, Linkedin's Constructor
func New(clientID, clientSecret, redirectURI string) *Linkedin {
	httpClient := internal.NewHTTPClient()
	return &Linkedin{
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
func (l *Linkedin) GetAuthURI() (string, error) {
	return "", nil
}

// GetAccessToken function
func (l *Linkedin) GetAccessToken(ctx context.Context, authorizationCode string) error {

	if l.ClientID == "" {
		return internal.NewErrorEmptyValue("client id")
	}

	if l.ClientSecret == "" {
		return internal.NewErrorEmptyValue("client secret")
	}

	if l.RedirectURI == "" {
		internal.NewErrorEmptyValue("redirect uri")
	}

	form := url.Values{}
	form.Add("grant_type", "authorization_code")
	form.Add("code", authorizationCode)
	form.Add("client_id", l.ClientID)
	form.Add("client_secret", l.ClientSecret)
	form.Add("redirect_uri", l.RedirectURI)

	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	var response struct {
		AccessToken      string `json:"access_token"`
		ExpiresIn        int    `json:"expires_in"`
		Error            string `json:"error"`
		ErrorDescription string `json:"error_description"`
	}

	err := l.httpClient.Execute(ctx, "POST", l.TokenURI, strings.NewReader(form.Encode()), &response, headers)

	if err != nil {
		return err
	}

	if len(response.Error) > 0 {
		return errors.New(response.ErrorDescription)
	}

	l.Token = response.AccessToken

	return nil
}

// GetUser function
func (l *Linkedin) GetUser(ctx context.Context) (api.Result, error) {

	if l.Token == "" {
		return nil, internal.NewErrorEmptyValue("access token")
	}

	headers := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", l.Token),
	}

	var response User

	uri := fmt.Sprintf("%s:(id,num-connections,picture-url,email-address,firstName,lastName,headline,siteStandardProfileRequest)?format=json", l.APIURI)

	err := l.httpClient.Execute(ctx, "GET", uri, nil, &response, headers)

	if err != nil {
		return nil, err
	}

	if len(response.Message) > 0 {
		return nil, errors.New(response.Message)
	}

	return &response, nil
}
