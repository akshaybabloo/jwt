package jwt

import (
	"encoding/base64"
	"encoding/json"
	"strings"
	"time"

	"github.com/akshaybabloo/jwt/providers"
)

// Token after an authentication is done
type Token struct {

	// JWT AccessToken
	AccessToken string `json:"access_token"`

	// JWT RefreshToken
	RefreshToken string `json:"refresh_token"`

	// JWT IdToken
	IdToken string `json:"id_token"`

	// ClientID of the provider
	ClientID string `json:"client_id"`
}

// Valid validates Token.AccessToken
func (t *Token) Valid() (bool, error) {
	var common providers.Common

	err := t.UnmarshalAccessToken(&common)
	if err != nil {
		return false, err
	}

	iat := UnixTimeToTime(common.Iat)
	exp := UnixTimeToTime(common.Exp)

	if !InTimeSpan(iat, exp, exp) {
		return false, TokenExpiredError
	}

	if common.Aud != t.ClientID {
		return false, ClientIDMismatchError
	}

	return true, nil
}

// UnmarshalAccessToken converts Token.AccessToken to string map
func (t *Token) UnmarshalAccessToken(v interface{}) error {
	token, err := t.DecodeToken(t.AccessToken)
	if err != nil {
		return err
	}
	err = t.toInterface(token, &v)
	if err != nil {
		return err
	}
	return nil
}

// UnmarshalRefreshToken converts Token.RefreshToken to string map
func (t *Token) UnmarshalRefreshToken(v interface{}) error {
	token, err := t.DecodeToken(t.RefreshToken)
	if err != nil {
		return err
	}
	err = t.toInterface(token, &v)
	if err != nil {
		return err
	}
	return nil
}

// UnmarshalIdToken converts Token.IdToken to string map
func (t *Token) UnmarshalIdToken(v interface{}) error {
	token, err := t.DecodeToken(t.IdToken)
	if err != nil {
		return err
	}
	err = t.toInterface(token, &v)
	if err != nil {
		return err
	}
	return nil
}

func (t *Token) DecodeToken(token string) ([]byte, error) {
	token = strings.Split(token, ".")[1]
	decodeString, err := base64.RawURLEncoding.DecodeString(token)
	if err != nil {
		return nil, err
	}
	return decodeString, nil
}

func (t *Token) toInterface(token []byte, s interface{}) error {
	err := json.Unmarshal(token, &s)
	if err != nil {
		return err
	}
	return nil
}

// UnixTimeToTime converts Unix time stamp to time.Time
func UnixTimeToTime(s int64) time.Time {
	return time.Unix(s, 0)
}

// InTimeSpan checks if two time stamps are in a given span
func InTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}
