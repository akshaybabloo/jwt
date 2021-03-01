package jwt

import "errors"

// TokenExpiredError occurs when the token is expired
var TokenExpiredError = errors.New("access token is expired")

// ClientIDMismatchError occurs when the token "aud" and Token.ClientID doesn't match
var ClientIDMismatchError = errors.New("client ID does not match token aud")

