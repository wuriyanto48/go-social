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

func newProvider(clientID, clientSecret, tenantID, redirectURI, scope string) *provider {
	providers := make(map[string]api.Service)
	providers["Facebook"] = facebook.New(clientID, clientSecret, redirectURI)
	providers["Google"] = google.New(clientID, clientSecret, redirectURI)
	providers["Github"] = github.New(clientID, clientSecret, redirectURI)
	providers["Linkedin"] = linkedin.New(clientID, clientSecret, redirectURI)
	providers["MicrosoftID"] = microsoftid.New(clientID, clientSecret, tenantID, redirectURI, scope)
	return &provider{providers}
}

// New , return api.Service implementation
func New(socialType Type, clientID, clientSecret, tenantID, redirectURI, scope string) (api.Service, error) {
	providers := newProvider(clientID, clientSecret, tenantID, redirectURI, scope)
	provider, ok := providers.providers[socialType.String()]

	if !ok {
		return nil, errors.New("invalid provider type")
	}

	return provider, nil
}
