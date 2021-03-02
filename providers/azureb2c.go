package providers

// AzureB2C can be used to unmarshall Azure AD B2C tokens
type AzureB2C struct {

	// The expiration time claim is the time at which the token becomes invalid,
	// represented in epoch time. Your app should use this claim to verify the
	// validity of the token lifetime.
	Exp int64 `json:"exp"`

	// This claim is the time at which the token becomes valid, represented in epoch
	// time. This is usually the same as the time the token was issued. Your app
	// should use this claim to verify the validity of the token lifetime.
	Nbf int64 `json:"nbf"`

	// The version of the ID token, as defined by Azure AD B2C.
	Ver string `json:"ver"`

	// This claim identifies the security token service (STS) that constructs and
	// returns the token. It also identifies the Azure AD directory in which the
	// user was authenticated. Your app should validate the issuer claim to ensure
	// that the token came from the v2.0 endpoint. It also should use the GUID
	// portion of the claim to restrict the set of tenants that can sign in to the
	// app.
	Iss string `json:"iss"`

	// This is the principal about which the token asserts information, such as the
	// user of an app. This value is immutable and cannot be reassigned or reused.
	// It can be used to perform authorization checks safely, such as when the token
	// is used to access a resource. By default, the subject claim is populated with
	// the object ID of the user in the directory. To learn more, see Azure Active
	// Directory B2C: Token, session, and single sign-on configuration.
	Sub string `json:"sub"`

	// An audience claim identifies the intended recipient of the token. For Azure
	// AD B2C, the audience is your app's Application ID, as assigned to your app in
	// the app registration portal. Your app should validate this value and reject
	// the token if it does not match.
	Aud string `json:"aud"`

	//	A nonce is a strategy used to mitigate token replay attacks. Your app can
	//	specify a nonce in an authorization request by using the nonce query
	//	parameter. The value you provide in the request will be emitted unmodified in
	//	the nonce claim of an ID token only. This allows your app to verify the value
	//	against the value it specified on the request, which associates the app's
	//	session with a given ID token. Your app should perform this validation during
	//	the ID token validation process.
	Nonce string `json:"nonce"`

	// The time at which the token was issued, represented in epoch time.
	Iat int64 `json:"iat"`

	// This claim is the time at which a user last entered credentials, represented in epoch time.
	AuthTime int64 `json:"auth_time"`

	// The immutable identifier for the user account in the tenant. It can be used
	// to perform authorization checks safely and as a key in database tables. This
	// ID uniquely identifies the user across applications - two different
	// applications signing in the same user will receive the same value in the oid
	// claim. This means that it can be used when making queries to Microsoft online
	// services, such as the Microsoft Graph. The Microsoft Graph will return this
	// ID as the id property for a given user account.
	Oid string `json:"oid"`

	// The user's given name (also known as first name).
	GivenName string `json:"given_name"`

	// The user's surname (also known as last name).
	FamilyName string `json:"family_name"`

	// This is the name of the policy that was used to acquire the token.
	Tfp string `json:"tfp"`

	// A code hash is included in an ID token only when the token is issued together
	// with an OAuth 2.0 authorization code. A code hash can be used to validate the
	// authenticity of an authorization code. For more details on how to perform
	// this validation, see the OpenID Connect specification.
	CHash string `json:"c_hash"`

	// An access token hash is included in an ID token only when the token is issued
	// together with an OAuth 2.0 access token. An access token hash can be used to
	// validate the authenticity of an access token. For more details on how to
	// perform this validation, see the OpenID Connect specification .
	AtHash string `json:"at_hash"`

	// Scope requested from the user
	Scp string `json:"scp"`

	// The providers access token. This will only be filled if you are using an external provider eg: GitHub.
	IdpAccessToken string `json:"idp_access_token"`

	// The identity provider used by the user to authenticate. This will only be
	// filled if you are using an external provider eg: GitHub.
	Idp string `json:"idp"`
}
