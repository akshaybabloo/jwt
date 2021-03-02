package providers

type AzureB2C struct {
	Exp        int64  `json:"exp"`
	Nbf        int64  `json:"nbf"`
	Ver        string `json:"ver"`
	Iss        string `json:"iss"`
	Sub        string `json:"sub"`
	Aud        string `json:"aud"`
	Nonce      string `json:"nonce"`
	Iat        int64  `json:"iat"`
	AuthTime   int64  `json:"auth_time"`
	Oid        string `json:"oid"`
	GivenName  string `json:"given_name"`
	FamilyName string `json:"family_name"`
	Tfp        string `json:"tfp"`
	CHash      string `json:"c_hash"`
	AtHash     string `json:"at_hash"`

	IdpAccessToken string `json:"idp_access_token"`
	Idp            string `json:"idp"`
}
