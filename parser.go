package jwt

import (
	"net/url"
)

// ParseURL gets the query from the URL and returns Token
func ParseURL(u, clientID string) (Token, error) {

	parse, err := url.Parse(u)
	if err != nil {
		return Token{}, err
	}

	return Token{
		AccessToken:  parse.Query().Get("access_token"),
		RefreshToken: parse.Query().Get("refresh_token"),
		IdToken:      parse.Query().Get("id_token"),
		ClientID:     clientID,
	}, nil
}

// ParseAccessToken parses access token with client ID to return Token
func ParseAccessToken(t, clientId string) Token {
	return Token{
		AccessToken: t,
		ClientID: clientId,
	}
}

// ParseIdToken parses ID token to return Token
func ParseIdToken(t string) Token {
	return Token{
		AccessToken: t,
	}
}

// ParseRefreshToken parses refresh token to return Token
func ParseRefreshToken(t string) Token {
	return Token{
		AccessToken: t,
	}
}
