package providers

// Common token properties
type Common struct {

	// The expiration time claim is the time at which the token becomes invalid,
	// represented in epoch time. Your app should use this claim to verify the
	// validity of the token lifetime.
	Exp      int64 `json:"exp"`

	// The time at which the token was issued, represented in epoch time.
	Iat      int64 `json:"iat"`

	// This claim is the time at which a user last entered credentials, represented in epoch time.
	AuthTime int64 `json:"auth_time"`

	// An audience claim identifies the intended recipient of the token.
	// Your app should validate this value and reject
	// the token if it does not match.
	Aud        string `json:"aud"`
}
