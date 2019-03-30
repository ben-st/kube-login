package cmd

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

// TokenSet needed for oidc token return
type TokenSet struct {
	IDToken      string
	RefreshToken string
}

// GetTokenSet returns id and refresh token from oidc provider
func (c *authConfig) GetTokenSet() *TokenSet {

	formData := url.Values{
		"client_id":     {c.ClientID},
		"client_secret": {c.ClientSecret},
		"grant_type":    {"password"},
		"scope":         {"openid profile email"},
		"username":      {c.User},
		"password":      {c.Password},
	}

	// if insecure for oidc is set we will skip tls verfification, use with caution!
	if c.insecureOIDC {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	resp, err := http.PostForm(c.URL, formData)
	if err != nil {
		log.Fatalln(err)
	}

	if resp.StatusCode != 200 {
		log.Fatalln("HTTP request failed with Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	// interface to string, struct expects strings
	idToken := fmt.Sprint(result["id_token"])
	refreshToken := fmt.Sprint(result["refresh_token"])

	return &TokenSet{
		IDToken:      idToken,
		RefreshToken: refreshToken,
	}

}
