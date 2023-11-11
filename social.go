package social

import (
	"errors"

	"github.com/wuriyanto48/go-social/api"
	"github.com/wuriyanto48/go-social/pkg/facebook"
	"github.com/wuriyanto48/go-social/pkg/github"
	"github.com/wuriyanto48/go-social/pkg/google"
	"github.com/wuriyanto48/go-social/pkg/linkedin"
	"github.com/wuriyanto48/go-social/pkg/microsoftid"
)

// Type generic type of social login
type Type int

// Provider structure
type provider struct {
	providers map[string]api.Service
}

const (
	// Facebook type
	Facebook Type = iota

	// Google Type
	Google

	// Linkedin Type
	Linkedin

	// Twitter Type
	Twitter

	// Github Type
	Github

	// MicrosoftID Type
	MicrosoftID
)

const (
	// DefaultHTTPClientTimeout for HTTP Client
	DefaultHTTPClientTimeout = 5
)

// String function
// returns string of Type
func (t Type) String() string {
	switch t {
	case Facebook:
		return "Facebook"
	case Google:
		return "Google"
	case Linkedin:
		return "Linkedin"
	case Twitter:
		return "Twitter"
	case Github:
		return "Github"
	case MicrosoftID:
		return "MicrosoftID"
	default:
		panic("unknown social media")
	}
}

func newProvider(clientID, clientSecret, tenantID, redirectURI, scope string, timeout int) *provider {
	providers := make(map[string]api.Service)
	providers["Facebook"] = facebook.New(clientID, clientSecret, redirectURI, timeout)
	providers["Google"] = google.New(clientID, clientSecret, redirectURI, timeout)
	providers["Github"] = github.New(clientID, clientSecret, redirectURI, timeout)
	providers["Linkedin"] = linkedin.New(clientID, clientSecret, redirectURI, timeout)
	providers["MicrosoftID"] = microsoftid.New(clientID, clientSecret, tenantID, redirectURI, scope, timeout)
	return &provider{providers}
}

// New , return api.Service implementation
func New(socialType Type, clientID, clientSecret, tenantID, redirectURI, scope string, timeout int) (api.Service, error) {
	// set default http client timeout to 5 seconds
	if timeout <= 0 {
		timeout = DefaultHTTPClientTimeout
	}

	providers := newProvider(clientID, clientSecret, tenantID, redirectURI, scope, timeout)
	provider, ok := providers.providers[socialType.String()]

	if !ok {
		return nil, errors.New("invalid provider type")
	}

	return provider, nil
}
