package social

import (
	"testing"

	"github.com/wuriyanto48/go-social/pkg/facebook"
)

func TestSocial(t *testing.T) {
	expected := "Google"

	googleType := Google

	if googleType.String() != expected {
		t.Error("Error Social type conversion")
	}
}

func TestProviderType(t *testing.T) {

	t.Run("Should return Facebook", func(t *testing.T) {
		expected1 := "https://graph.facebook.com/oauth/access_token"
		expected2 := "https://graph.facebook.com/v2/oauth/access_token"

		result, err := New(Facebook, "client_id", "client_secret", "http://localhost:8080/callback")

		if err != nil {
			t.Error(err)
		}

		// convert result to facebook
		actual, _ := result.(*facebook.Facebook)

		if actual.TokenURI != expected1 {
			t.Error("value is not Facebook")
		}

		// test mutate
		actual.TokenURI = "https://graph.facebook.com/v2/oauth/access_token"

		if actual.TokenURI != expected2 {
			t.Error("value is not Facebook")
		}

	})
}
