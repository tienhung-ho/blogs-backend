package authmodel

type UserToken struct {
	AccessToken  string
	RefreshToken string
}

func NewUserToken(accessToken, refreshToken string) *UserToken {
	return &UserToken{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}
