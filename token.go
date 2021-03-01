package jwt

import (
	"encoding/base64"
	"encoding/json"
	"strconv"
	"strings"
	"time"
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

	token, err := t.UnmarshalAccessToken()
	if err != nil {
		return false, err
	}

	iat, err := UnixTimeToTime(token["iat"].(string))
	if err != nil {
		return false, err
	}
	exp, err := UnixTimeToTime(token["exp"].(string))
	if err != nil {
		return false, err
	}

	if !InTimeSpan(iat, exp, exp) {
		return false, TokenExpiredError
	}

	if token["aud"] != t.ClientID {
		return false, ClientIDMismatchError
	}

	return true, nil
}

// UnmarshalAccessToken converts Token.AccessToken to string map
func (t *Token) UnmarshalAccessToken() (map[string]interface{}, error) {
	token, err := t.DecodeToken(t.AccessToken)
	if err != nil {
		return nil, err
	}
	toInterface, err := t.toInterface(token)
	if err != nil {
		return nil, err
	}
	return toInterface, nil
}

// UnmarshalRefreshToken converts Token.RefreshToken to string map
func (t *Token) UnmarshalRefreshToken() (map[string]interface{}, error) {
	token, err := t.DecodeToken(t.RefreshToken)
	if err != nil {
		return nil, err
	}
	toInterface, err := t.toInterface(token)
	if err != nil {
		return nil, err
	}
	return toInterface, nil
}

// UnmarshalIdToken converts Token.IdToken to string map
func (t *Token) UnmarshalIdToken() (map[string]interface{}, error) {
	token, err := t.DecodeToken(t.IdToken)
	if err != nil {
		return nil, err
	}
	toInterface, err := t.toInterface(token)
	if err != nil {
		return nil, err
	}
	return toInterface, nil
}

func (t *Token) DecodeToken(token string) ([]byte, error) {
	token = strings.Split(token, ".")[1]
	decodeString, err := base64.StdEncoding.WithPadding(base64.NoPadding).DecodeString(token)
	if err != nil {
		return nil, err
	}
	return decodeString, nil
}

func (t *Token) toInterface(token []byte) (map[string]interface{}, error) {
	var data map[string]interface{}
	err := json.Unmarshal(token, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// UnixTimeToTime converts Unix time stamp to time.Time
func UnixTimeToTime(s string) (time.Time, error) {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return time.Time{}, err
	}
	return time.Unix(i, 0), nil
}

// InTimeSpan checks if two time stamps are in a given span
func InTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}
