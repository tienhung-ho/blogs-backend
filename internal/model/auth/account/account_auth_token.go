package accountautmodel

type AccountToken struct {
	AccessToken  string
	RefreshToken string
}

func NewAccountToken(accessToken, refreshToken string) *AccountToken {
	return &AccountToken{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}
