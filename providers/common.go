package providers

type Common struct {
	Exp      int64 `json:"exp"`
	Iat      int64 `json:"iat"`
	AuthTime int64 `json:"auth_time"`
	Aud        string `json:"aud"`
}
