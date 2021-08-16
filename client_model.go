package oppo

type Ret struct {
	Errno int64 `json:"errno"`
}

type RefreshTokenParams struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type RefreshTokenRes struct {
	Ret
	Data TokenInfo `json:"data"`
}

type TokenInfo struct {
	AccessToken string `json:"access_token"`
	ExpireIn    int64  `json:"expire_in"`
}
