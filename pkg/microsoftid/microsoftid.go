package microsoftid

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
	// DefaultAuthURI default Authorization URI for Microsoft Identity Platform
	DefaultAuthURI = "https://login.microsoftonline.com/%s/oauth2/v2.0/authorize?client_id=%s&response_type=code&scope=%s&redirect_uri=%s"

	// DefaultTokenURI default Token URI for Microsoft Identity Platform
	DefaultTokenURI = "https://login.microsoftonline.com/%s/oauth2/v2.0/token"

	// DefaultAPIRURI default API URI for Microsoft Identity Platform
	DefaultAPIRURI = "https://graph.microsoft.com/v1.0/me"
)

// MicrosoftID struct
type MicrosoftID struct {
	TenantID     string
	ClientID     string
	ClientSecret string
	AuthURI      string
	APIURI       string
	TokenURI     string
	RedirectURI  string
	Scope        string
	Token        string
	httpClient   *pkg.HTTPClient
}

// New function, MicrosoftID's Constructor
func New(clientID, clientSecret, tenantID, redirectURI, scope string) *MicrosoftID {
	httpClient := pkg.NewHTTPClient()
	return &MicrosoftID{
		TenantID:     tenantID,
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURI:  redirectURI,
		Scope:        scope,
		httpClient:   httpClient,
		AuthURI:      DefaultAuthURI,
		APIURI:       DefaultAPIRURI,
		TokenURI:     DefaultTokenURI,
	}
}

// GetAuthURI function
func (g *MicrosoftID) GetAuthURI() (string, error) {
	return fmt.Sprintf(g.AuthURI, g.TenantID, g.ClientID, g.Scope, g.RedirectURI), nil
}

// GetAccessToken function
func (g *MicrosoftID) GetAccessToken(ctx context.Context, authorizationCode string) error {

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
	}

	var response struct {
		AccessToken      string `json:"access_token"`
		ExpiresIn        int    `json:"expires_in"`
		ExtExpiresIn     int    `json:"ext_expires_in"`
		RefreshToken     string `json:"refresh_token"`
		Scope            string `json:"scope"`
		TokenType        string `json:"token_type"`
		IDToken          string `json:"id_token"`
		Error            string `json:"error"`
		ErrorDescription string `json:"error_description"`
		ErrorCodes       []int  `json:"error_codes"`
		Timestamp        string `json:"timestamp"`
		TraceID          string `json:"trace_id"`
		CorrelationID    string `json:"correlation_id"`
		ErrorURI         string `json:"error_uri"`
	}

	g.TokenURI = fmt.Sprintf(g.TokenURI, g.TenantID)

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
func (g *MicrosoftID) GetUser(ctx context.Context) (api.Result, error) {

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

	if response.Error != nil {
		return nil, errors.New(response.Error.Message)
	}

	return &response, nil
}
