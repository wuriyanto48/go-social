package social

import (
	"errors"

	"github.com/wuriyanto48/go-social/api"
	"github.com/wuriyanto48/go-social/internal/facebook"
	"github.com/wuriyanto48/go-social/internal/github"
	"github.com/wuriyanto48/go-social/internal/google"
	"github.com/wuriyanto48/go-social/internal/linkedin"
)

// Type generic type of social login
type Type int

// Provider structure
type Provider struct {
	Providers map[string]api.Service
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
	default:
		return "Facebook"
	}
}

func newProvider(clientID, clientSecret, redirectURI string) *Provider {
	providers := make(map[string]api.Service)
	providers["Facebook"] = facebook.New(clientID, clientSecret, redirectURI)
	providers["Google"] = google.New(clientID, clientSecret, redirectURI)
	providers["Github"] = github.New(clientID, clientSecret, redirectURI)
	providers["Linkedin"] = linkedin.New(clientID, clientSecret, redirectURI)
	return &Provider{providers}
}

// New , return api.Service implementation
func New(socialType Type, clientID, clientSecret, redirectURI string) (api.Service, error) {
	providers := newProvider(clientID, clientSecret, redirectURI)
	provider, ok := providers.Providers[socialType.String()]

	if !ok {
		return nil, errors.New("invalid provider type")
	}

	return provider, nil
}
